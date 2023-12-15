package logger

import "log"

type Logger struct{}

func (logger *Logger) Debug(message string) {
	log.Print(message)
}

func (logger *Logger) Error(message string) {
	log.Fatal(message)
}

func (logger *Logger) Info(message string) {
	log.Print(message)
}
