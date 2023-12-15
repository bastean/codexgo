package models

type Logger interface {
	Debug(message string)
	Error(message string)
	Info(message string)
}
