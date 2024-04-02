package communication

import (
	"bytes"
	"context"
	"fmt"
	"net/smtp"
	"os"

	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/domain/template"
)

type Smtp struct {
	smtp.Auth
	MIMEHeaders             string
	URL, Username, Password string
}

func (client *Smtp) Send(mailTemplate model.MailTemplate) {
	switch mail := mailTemplate.(type) {
	case *template.AccountConfirmationMail:
		client.SendAccountConfirmation(mail)
	}
}

func (client *Smtp) SendMail(to []string, message []byte) error {
	return smtp.SendMail(client.URL, client.Auth, client.Username, to, message)
}

func (client *Smtp) SendAccountConfirmation(mail *template.AccountConfirmationMail) {
	var message bytes.Buffer

	// TODO!: mail.To[0]
	headers := fmt.Sprintf("From: %s\n"+"To: %s\n"+"Subject: Account Confirmation", client.Username, mail.To[0])

	message.Write([]byte(fmt.Sprintf("%s\n%s\n", headers, client.MIMEHeaders)))

	mail.ConfirmationLink = fmt.Sprintf("%s/verify/%s", os.Getenv("URL"), mail.ConfirmationLink)

	AccountConfirmation(mail.Username, mail.ConfirmationLink).Render(context.Background(), &message)

	err := client.SendMail(mail.To, message.Bytes())

	if err != nil {
		panic(err)
	}
}

func NewNotifySmtpMail(host, port, username, password string) model.Mail {
	return &Smtp{
		Auth:        smtp.PlainAuth("", username, password, host),
		MIMEHeaders: "MIME-version: 1.0;\n" + "Content-Type: text/html; charset=\"UTF-8\";\n\n",
		URL:         host + ":" + port,
		Username:    username,
		Password:    password,
	}
}
