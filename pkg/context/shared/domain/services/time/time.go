package time

import (
	"log"
	"os"
	"strconv"
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
		date := Time{time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)}

		switch {
		case os.Getenv("GOTEST_FROZEN_BEFORE") != "":
			value, err := strconv.Atoi(os.Getenv("GOTEST_FROZEN_BEFORE"))

			if err != nil {
				log.Panic(err)
			}

			date = date.Add(-Duration(value))
		case os.Getenv("GOTEST_FROZEN_AFTER") != "":
			value, err := strconv.Atoi(os.Getenv("GOTEST_FROZEN_AFTER"))

			if err != nil {
				log.Panic(err)
			}

			date = date.Add(Duration(value))
		}

		return date
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
