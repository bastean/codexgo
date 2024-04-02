package model

type Mail interface {
	Send(template MailTemplate)
}
