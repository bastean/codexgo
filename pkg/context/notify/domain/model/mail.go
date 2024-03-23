package model

type Mail interface {
	Send(to []string, msg string)
}
