package logger

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/github"
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

func (l *Logger) NewEntry(event github.PushEvent, eventtype string, name string) *LogEntry {
	ts := time.Now()

	le := &LogEntry{
		ID: fmt.Sprintf("%d", ts.Unix()),
		StartTime: Timestamp{
			ts,
		},
		Name: name,
		Event: Event{
			ID:          *event.After,
			Type:        eventtype,
			Description: *event.HeadCommit.Message,
			Sender: EventSender{
				Name:       *event.Sender.Login,
				AvatarURL:  *event.Sender.AvatarURL,
				ProfileURL: *event.Sender.HTMLURL,
			},
		},
	}

	le.writer = &Writer{
		prefix: le.GetID(),
		redis:  l.redis,
	}

	return le
}

func (l *Logger) CreateEntry(e *LogEntry) {
	key := fmt.Sprintf("%s:builds", e.Name)
	l.redis.ZAdd(key, redis.Z{Score: float64(e.StartTime.Time.Unix()), Member: e.GetID()}).Result()
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

func (l *Logger) GetLog(host, namespace, name, commit, timestamp string) string {
	key := fmt.Sprintf("%s:%s:%s:%s:%s:log", host, namespace, name, timestamp, commit)
	data, _ := l.redis.Get(key).Result()
	return data
}
