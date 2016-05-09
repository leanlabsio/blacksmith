package logger

import (
	"time"
)

type LogEntry struct {
	Name      string `json:"name"`
	StartTime int64  `json:"start_time"`
	event     interface{}
	Duration  time.Duration `json:"duration"`
	writer    *Writer
}

func (e *LogEntry) Close() error {
	d := time.Since(time.Unix(e.StartTime, 0))
	e.Duration = d
	e.writer.WriteEntry(e)
	return nil
}

func (e *LogEntry) Write(p []byte) (int, error) {
	return e.writer.Write(p)
}
