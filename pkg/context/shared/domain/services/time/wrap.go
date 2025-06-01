package time

import (
	"time"
)

const (
	RFC3339Nano = time.RFC3339Nano
)

const (
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
)

const (
	Day   = Hour * 24
	Week  = Day * 7
	Month = Week * 4
	Year  = Month * 12
)

var (
	Sleep = time.Sleep
)

type (
	Duration = time.Duration
)
