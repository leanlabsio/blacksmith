package logger

import (
	"encoding/json"
	"fmt"
	"gopkg.in/redis.v3"
	"time"
)

type Logger struct {
	redis *redis.Client
}

func New(r *redis.Client) *Logger {
	return &Logger{
		redis: r,
	}
}

func (l *Logger) CreateEntry(name string) *LogEntry {
	ts := time.Now()
	unixTs := ts.Unix()
	buildentry := fmt.Sprintf("%s:%d", name, unixTs)
	key := fmt.Sprintf("%s:builds", name)

	l.redis.ZAdd(key, redis.Z{Score: float64(unixTs), Member: buildentry}).Result()

	writer := &Writer{
		redis:  l.redis,
		prefix: buildentry,
	}
	le := &LogEntry{
		writer: writer,
		StartTime: Timestamp{
			ts,
		},
		Name: buildentry,
	}
	writer.WriteEntry(le)

	return le
}

func (l *Logger) ListEntries(host, namespace, name string) []LogEntry {
	key := fmt.Sprintf("%s:%s:%s:builds", host, namespace, name)

	data, _ := l.redis.ZRevRangeByScoreWithScores(key, redis.ZRangeByScore{Min: "-inf", Max: "+inf", Offset: 0, Count: 50}).Result()

	var builds []LogEntry

	for _, item := range data {
		entryKey := item.Member.(string)
		e, _ := l.redis.Get(entryKey).Result()

		var entry LogEntry
		json.Unmarshal([]byte(e), &entry)
		builds = append(builds, entry)
	}

	return builds
}
