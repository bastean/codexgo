package smtp

import (
	"fmt"
	"net/smtp"
)

type SMTP struct {
	smtp.Auth
	SMTPServerURL, Username, Password, ServerURL string
}

func (*SMTP) SetHeader(key, value string) string {
	return fmt.Sprintf("%s: %s\n", key, value)
}

func (client *SMTP) Headers(to, subject string) string {
	header := client.SetHeader("From", client.Username)

	header += client.SetHeader("To", to)

	header += client.SetHeader("Subject", subject)

	header += client.SetHeader("MIME-version", "1.0")

	header += client.SetHeader("Content-Type", "text/html; charset=\"UTF-8\"")

	header += "\n"

	return header
}

func (client *SMTP) SendMail(to []string, message []byte) error {
	return smtp.SendMail(client.SMTPServerURL, client.Auth, client.Username, to, message)
}

func Open(host, port, username, password, serverURL string) *SMTP {
	return &SMTP{
		Auth:          smtp.PlainAuth("", username, password, host),
		SMTPServerURL: fmt.Sprintf("%s:%s", host, port),
		Username:      username,
		Password:      password,
		ServerURL:     serverURL,
	}
}
