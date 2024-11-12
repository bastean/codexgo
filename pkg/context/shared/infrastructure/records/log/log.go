package log

import (
	"log"

	"github.com/fatih/color"
)

type Log struct {
	Cyan, Red, Blue, Green *color.Color
}

func (l *Log) Debug(message string) {
	log.Println(l.Cyan.Sprint(message))
}

func (l *Log) Error(message string) {
	log.Println(l.Red.Sprint(message))
}

func (l *Log) Fatal(message string) {
	log.Fatalln(l.Red.Sprint(message))
}

func (l *Log) Info(message string) {
	log.Println(l.Blue.Sprint(message))
}

func (l *Log) Success(message string) {
	log.Println(l.Green.Sprint(message))
}

func New() *Log {
	return &Log{
		Cyan:  color.New(color.FgCyan, color.Bold),
		Red:   color.New(color.FgRed, color.Bold),
		Blue:  color.New(color.FgBlue, color.Bold),
		Green: color.New(color.FgGreen, color.Bold),
	}
}
