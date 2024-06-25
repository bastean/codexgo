package logger

import (
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/loggers"
)

var (
	Logger = new(loggers.Logger)
	Debug  = Logger.Debug
	Error  = Logger.Error
	Fatal  = Logger.Fatal
	Info   = Logger.Info
)
