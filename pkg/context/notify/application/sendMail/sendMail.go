package sendMail

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
)

type SendMail struct {
	model.Mail
}

func (sendMail *SendMail) Run(mail model.MailTemplate) {
	sendMail.Mail.Send(mail)
}

func NewSendMail(mail model.Mail) *SendMail {
	return &SendMail{
		Mail: mail,
	}
}
