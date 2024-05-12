package logger

import (
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/slogger"
)

var Logger = new(slogger.Logger)

var Debug = Logger.Debug
var Error = Logger.Error
var Fatal = Logger.Fatal
var Info = Logger.Info
