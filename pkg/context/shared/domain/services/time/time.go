package time

import (
	"log"
	"os"
	"time"
)

type Time struct {
	time time.Time
}

func (t Time) IsZero() bool {
	return t.time.IsZero()
}

func (t Time) UTC() Time {
	return Time{t.time.UTC()}
}

func (t Time) Unix() int64 {
	return t.time.Unix()
}

func (t Time) Format() string {
	return t.time.UTC().Format(RFC3339Nano)
}

func (t Time) Add(d Duration) Time {
	return Time{t.time.Add(d)}
}

func (t Time) Before(u Time) bool {
	return t.time.Before(u.time)
}

func (t Time) Equal(u Time) bool {
	return t.time.Equal(u.time)
}

func (t Time) After(u Time) bool {
	return t.time.After(u.time)
}

func Now() Time {
	if _, ok := os.LookupEnv("GOTEST_FROZEN"); ok {
		return Time{time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)}
	}

	return Time{time.Now().UTC()}
}

func Parse(value string) Time {
	time, err := time.Parse(RFC3339Nano, value)

	if err != nil {
		log.Panic("Time format is not valid")
	}

	return Time{time}
}
