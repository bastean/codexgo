package log

import (
	"fmt"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"

	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
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
)

func Logo() {
	figureCodex := figure.NewFigure("codex", "speed", true).Slicify()
	figureGo := figure.NewFigure("GO", "speed", true).Slicify()

	var width, fixedWidth int

	for _, line := range figureCodex {
		width = len(line)

		if width > fixedWidth {
			fixedWidth = width
		}
	}

	for i, line := range figureCodex {
		width = len(line)

		if width < fixedWidth {
			line += strings.Repeat(" ", (fixedWidth - width))
		}

		fmt.Println(White(line), Cyan(figureGo[i]))
	}

	fmt.Println()
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
