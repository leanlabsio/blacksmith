package logger

import (
	"time"
)

type LogEntry struct {
	name      string
	startTime int64
	event     interface{}
	duration  time.Duration
	writer    *Writer
}

func (e *LogEntry) Start() {
	e.writer.WriteEntry(e)
}

func (e *LogEntry) Finish() {
	e.writer.WriteEntry(e)
}

func (e *LogEntry) Write(p []byte) (int, error) {
	return e.writer.Write(p)
}
