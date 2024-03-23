package sendMail

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
)

type SendMail struct {
	model.Mail
}

func (sendMail *SendMail) Run(to, msg string) {
	sendMail.Mail.Send([]string{to}, msg)
}

func NewSendMail(mail model.Mail) *SendMail {
	return &SendMail{
		Mail: mail,
	}
}
