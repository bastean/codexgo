package log

import (
	"log"

	"github.com/fatih/color"
)

type Log struct {
	*log.Logger
	Cyan, Red, Blue, Green *color.Color
}

func (l *Log) Debug(format string, values ...any) {
	l.Println(l.Cyan.Sprintf(format, values...))
}

func (l *Log) Error(format string, values ...any) {
	l.Println(l.Red.Sprintf(format, values...))
}

func (l *Log) Fatal(format string, values ...any) {
	l.Fatalln(l.Red.Sprintf(format, values...))
}

func (l *Log) Info(format string, values ...any) {
	l.Println(l.Blue.Sprintf(format, values...))
}

func (l *Log) Success(format string, values ...any) {
	l.Println(l.Green.Sprintf(format, values...))
}

func New() *Log {
	return &Log{
		Logger: log.Default(),
		Cyan:   color.New(color.FgCyan, color.Bold),
		Red:    color.New(color.FgRed, color.Bold),
		Blue:   color.New(color.FgBlue, color.Bold),
		Green:  color.New(color.FgGreen, color.Bold),
	}
}
