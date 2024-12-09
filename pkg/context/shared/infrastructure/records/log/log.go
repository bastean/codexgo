package log

import (
	"log"

	"github.com/fatih/color"
)

type Log struct {
	*log.Logger
	Cyan, Red, Blue, Green *color.Color
}

func (l *Log) Debug(message string) {
	l.Println(l.Cyan.Sprint(message))
}

func (l *Log) Error(message string) {
	l.Println(l.Red.Sprint(message))
}

func (l *Log) Fatal(message string) {
	l.Fatalln(l.Red.Sprint(message))
}

func (l *Log) Info(message string) {
	l.Println(l.Blue.Sprint(message))
}

func (l *Log) Success(message string) {
	l.Println(l.Green.Sprint(message))
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
