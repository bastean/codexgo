package communication

import (
	"net/smtp"
	"os"

	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
)

var host = os.Getenv("SMTP_HOST")
var port = os.Getenv("SMTP_PORT")
var username = os.Getenv("SMTP_USERNAME")
var password = os.Getenv("SMTP_PASSWORD")

type Smtp struct {
	Auth smtp.Auth
}

func (client *Smtp) Send(to []string, msg string) {
	err := smtp.SendMail(host+":"+port, client.Auth, username, to, []byte(msg))

	if err != nil {
		panic(err)
	}
}

func NewNotifySmtpMail() model.Mail {
	return &Smtp{
		Auth: smtp.PlainAuth("", username, password, host),
	}
}
