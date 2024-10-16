package log

import (
	"log"

	"github.com/fatih/color"
)

type Log struct {
	Cyan, Red, Blue, Green *color.Color
}

func (color *Log) Debug(message string) {
	log.Println(color.Cyan.Sprint(message))
}

func (color *Log) Error(message string) {
	log.Println(color.Red.Sprint(message))
}

func (color *Log) Fatal(message string) {
	log.Fatalln(color.Red.Sprint(message))
}

func (color *Log) Info(message string) {
	log.Println(color.Blue.Sprint(message))
}

func (color *Log) Success(message string) {
	log.Println(color.Green.Sprint(message))
}

func New() *Log {
	return &Log{
		Cyan:  color.New(color.FgCyan, color.Bold),
		Red:   color.New(color.FgRed, color.Bold),
		Blue:  color.New(color.FgBlue, color.Bold),
		Green: color.New(color.FgGreen, color.Bold),
	}
}
