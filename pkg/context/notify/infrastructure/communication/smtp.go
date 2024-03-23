package communication

import (
	"log"

	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
)

type Smtp struct{}

func (smtp *Smtp) Send(to []string, msg string) {
	// TODO
	log.Println("TO:", to)
	log.Println("Msg:", msg)
}

func NewNotifySmtpMail() model.Mail {
	return new(Smtp)
}
