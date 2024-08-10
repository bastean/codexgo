package smtp

import (
	"fmt"
	"net/smtp"
)

type Auth struct {
	Host, Port, Username, Password string
}

type SMTP struct {
	smtp.Auth
	ServerURL, Username string
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
	return smtp.SendMail(
		client.ServerURL,
		client.Auth,
		client.Username,
		to,
		message,
	)
}

func Open(auth *Auth) *SMTP {
	return &SMTP{
		Auth:      smtp.PlainAuth("", auth.Username, auth.Password, auth.Host),
		ServerURL: fmt.Sprintf("%s:%s", auth.Host, auth.Port),
		Username:  auth.Username,
	}
}
