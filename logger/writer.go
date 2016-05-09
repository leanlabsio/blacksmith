package logger

import (
	"encoding/json"
	"gopkg.in/redis.v3"
)

type Writer struct {
	prefix string
	redis  *redis.Client
}

func (w *Writer) WriteEntry(e *LogEntry) {
	data, _ := json.Marshal(e)
	w.redis.Set(e.Name, data, 0)
}

func (w *Writer) Write(p []byte) (int, error) {
	w.redis.Append(w.prefix+":log", string(p)).Result()
	return len(p), nil
}
