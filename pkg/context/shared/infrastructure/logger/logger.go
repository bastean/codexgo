package logger

import (
	"log"

	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
)

type Logger struct{}

func (logger *Logger) Debug(message string) {
	log.Print(message)
}

func (logger *Logger) Error(message string) {
	log.Print(message)
}

func (logger *Logger) Fatal(message string) {
	log.Fatal(message)
}

func (logger *Logger) Info(message string) {
	log.Print(message)
}

func NewLogger() model.Logger {
	return new(Logger)
}
