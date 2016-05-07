package logger

import (
	"gopkg.in/redis.v3"
)

func New(name string, r *redis.Client) *Writer {
	return &Writer{
		name:  name,
		redis: r,
	}
}

type Writer struct {
	name  string
	redis *redis.Client
}

func (w *Writer) Write(p []byte) (int, error) {
	w.redis.Append(w.name, string(p)).Result()
	return len(p), nil
}
