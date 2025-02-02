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
	return fmt.Sprintf("Service:%s", service)
}

func Server(app string) string {
	return fmt.Sprintf("Server:%s", app)
}

func Starting(service string) {
	Info(fmt.Sprintf("Starting %s...", service))
}

func Started(service string) {
	Success(fmt.Sprintf("%s started", service))
}

func CannotBeStarted(service string) {
	Error(fmt.Sprintf("%s cannot be started", service))
}

func Stopping(service string) {
	Info(fmt.Sprintf("Stopping %s...", service))
}

func Stopped(service string) {
	Success(fmt.Sprintf("%s stopped", service))
}

func CannotBeStopped(service string) {
	Error(fmt.Sprintf("%s cannot be stopped", service))
}

func EstablishingConnectionWith(service string) {
	Info(fmt.Sprintf("Establishing connection with %s...", service))
}

func ConnectionEstablishedWith(service string) {
	Success(fmt.Sprintf("Connection established with %s", service))
}

func ConnectionFailedWith(service string) {
	Error(fmt.Sprintf("Connection failed with %s", service))
}

func ClosingConnectionWith(service string) {
	Info(fmt.Sprintf("Closing connection with %s...", service))
}

func ConnectionClosedWith(service string) {
	Success(fmt.Sprintf("Connection closed with %s", service))
}

func DisconnectionFailedWith(service string) {
	Error(fmt.Sprintf("Disconnection failed with %s", service))
}
