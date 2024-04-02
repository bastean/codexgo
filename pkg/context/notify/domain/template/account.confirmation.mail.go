package template

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
)

type AccountConfirmationMail struct {
	*Mail
	Username         string
	ConfirmationLink string
}

func NewAccountConfirmationMail(mail *Mail, username, confirmationLink string) model.MailTemplate {
	return &AccountConfirmationMail{
		Mail:             mail,
		Username:         username,
		ConfirmationLink: confirmationLink,
	}
}
