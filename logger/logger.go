package logger

import (
	"fmt"
	"gopkg.in/redis.v3"
	"time"
)

func New(path string, r *redis.Client) *Writer {
	timestamp := time.Now().Unix()
	key := fmt.Sprintf("%s:builds", path)

	buildentry := fmt.Sprintf("%s:%d:build", path, timestamp)
	logentry := fmt.Sprintf("%s:%d:log", path, timestamp)

	r.HMSet(buildentry, "user_name", "qwerty", "commit", "qwerty", "timestamp", string(timestamp)).Result()
	r.ZAdd(key, redis.Z{Score: float64(timestamp), Member: buildentry}).Result()

	return &Writer{
		name:  logentry,
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
