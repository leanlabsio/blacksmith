package logger

import (
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
	timestamp := time.Now().Unix()
	buildentry := fmt.Sprintf("%s:%d", name, timestamp)

	key := fmt.Sprintf("%s:builds", name)

	l.redis.ZAdd(key, redis.Z{Score: float64(timestamp), Member: buildentry}).Result()

	writer := &Writer{redis: l.redis, prefix: buildentry}
	le := &LogEntry{writer: writer, startTime: timestamp, name: buildentry}

	return le
}

func (l *Logger) ListEntries(name string) {}
