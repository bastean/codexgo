package model

type Logger interface {
	Debug(message string)
	Error(message string)
	Fatal(message string)
	Info(message string)
}
