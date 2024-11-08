package services

import (
	"time"
)

func TimeNow() string {
	return FormatTime(time.Now().UTC())
}

func FormatTime(value time.Time) string {
	return value.Format(time.RFC3339Nano)
}
