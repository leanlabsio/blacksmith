package logger

import (
	"strconv"
	"time"
)

type Timestamp struct {
	time.Time
}

func (t *Timestamp) MarshalJSON() ([]byte, error) {
	ts := strconv.Itoa(int(t.Time.Unix()))
	return []byte(ts), nil
}

func (t *Timestamp) UnmarshalJSON(b []byte) error {
	ts, err := strconv.Atoi(string(b))

	if err != nil {
		return err
	}

	t.Time = time.Unix(int64(ts), 0)
	return nil
}

type LogEntry struct {
	Name      string    `json:"name"`
	StartTime Timestamp `json:"start_time"`
	event     interface{}
	Duration  time.Duration `json:"duration"`
	writer    *Writer
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
