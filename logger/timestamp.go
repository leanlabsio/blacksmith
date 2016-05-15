package logger

import (
	"strconv"
	"time"
)

// Timestamp is a wrapper struct for time.Time,
// used to implement json.Marshaler and json.Unmarshaler interfaces,
// it's json representation is always unix timestamp.
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
