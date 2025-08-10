package log

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/array"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/ascii"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
)

const (
	FontName   = "speed"
	FontHeight = 5
)

var (
	Log     = log.New()
	Debug   = Log.Debug
	Error   = Log.Error
	Fatal   = Log.Fatal
	Info    = Log.Info
	Success = Log.Success
)

var (
	White = color.New(color.FgWhite, color.Bold).Sprint
	Cyan  = color.New(color.FgCyan, color.Bold).Sprint
	Black = color.New(color.FgBlack, color.Bold).Sprint
)

func Logo(version ...string) {
	figureCodex := figure.NewFigure("codex", FontName, true).Slicify()
	figureGo := figure.NewFigure("GO", FontName, true).Slicify()
	figureVersion := make([]string, FontHeight)

	latest, exists := array.Slice(version, 0)

	if exists {
		figureVersion[FontHeight-1] = "v" + latest
	}

	ascii.FixWidth(figureCodex, figureGo)

	for i := range FontHeight {
		fmt.Println(White(figureCodex[i]), Cyan(figureGo[i], Black(figureVersion[i])))
	}

	println()
}

func Service(service string) string {
	return "Service:" + service
}

func Server(app string) string {
	return "Server:" + app
}

func Starting(service string) {
	Info("Starting " + service + "...")
}

func Started(service string) {
	Success(service + " started")
}

func CannotBeStarted(service string) {
	Error(service + " cannot be started")
}

func Stopping(service string) {
	Info("Stopping " + service + "...")
}

func Stopped(service string) {
	Success(service + " stopped")
}

func CannotBeStopped(service string) {
	Error(service + " cannot be stopped")
}

func EstablishingConnectionWith(service string) {
	Info("Establishing connection with " + service + "...")
}

func ConnectionEstablishedWith(service string) {
	Success("Connection established with " + service)
}

func ConnectionFailedWith(service string) {
	Error("Connection failed with " + service)
}

func ClosingConnectionWith(service string) {
	Info("Closing connection with " + service + "...")
}

func ConnectionClosedWith(service string) {
	Success("Connection closed with " + service)
}

func DisconnectionFailedWith(service string) {
	Error("Disconnection failed with " + service)
}
