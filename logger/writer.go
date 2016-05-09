package logger

import (
	"gopkg.in/redis.v3"
)

type Writer struct {
	prefix string
	redis  *redis.Client
}

func (w *Writer) WriteEntry(e *LogEntry) {
	w.redis.HMSet(e.name, "user_name", "qwerty", "commit", "qwerty").Result()
}

func (w *Writer) Write(p []byte) (int, error) {
	w.redis.Append(w.prefix+":log", string(p)).Result()
	return len(p), nil
}
