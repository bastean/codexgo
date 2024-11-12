package services

import (
	"time"
)

func FormatTime(value time.Time) string {
	return value.Format(time.RFC3339Nano)
}

func TimeNow() string {
	return FormatTime(time.Now().UTC())
}
