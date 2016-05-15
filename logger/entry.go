package logger

import (
	"fmt"
	"time"
)

type LogEntry struct {
	ID        string        `json:"id"`
	Name      string        `json:"name"`
	StartTime Timestamp     `json:"start_time"`
	Duration  time.Duration `json:"duration"`
	Event     Event         `json:"event"`
	State     string        `json:"state"`

	writer *Writer
}

type Event struct {
	ID          string      `json:"id"`
	Type        string      `json:"type"`
	Sender      EventSender `json:"sender"`
	Description string      `json:"description"`
}

type EventSender struct {
	Name       string `json:"name"`
	AvatarURL  string `json:"avatar_url"`
	ProfileURL string `json:"profile_url"`
}

func (e *LogEntry) GetID() string {
	return fmt.Sprintf("%s:%s:%s", e.Name, e.ID, e.Event.ID)
}

func (e *LogEntry) Start() error {
	e.writer.WriteEntry(e)
	return nil
}

func (e *LogEntry) Close() error {
	d := time.Since(e.StartTime.Time)
	e.Duration = d
	e.writer.WriteEntry(e)
	return nil
}

func (e *LogEntry) Write(p []byte) (int, error) {
	return e.writer.Write(p)
}
