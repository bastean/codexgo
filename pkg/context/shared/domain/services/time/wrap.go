package time

import (
	"time"
)

const (
	RFC3339Nano = time.RFC3339Nano
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
)

var (
	Sleep = time.Sleep
)

type (
	Duration = time.Duration
)
