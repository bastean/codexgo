package smtp

import (
	"fmt"
	"net/smtp"
)

type SMTP struct {
	smtp.Auth
	MIMEHeaders, SMTPServerURL, Username, Password, ServerURL string
}

func (client *SMTP) Headers(to, subject string) string {
	return fmt.Sprintf("From: %s\n"+"To: %s\n"+"Subject: %s\n"+"%s\n", client.Username, to, subject, client.MIMEHeaders)
}

func (client *SMTP) SendMail(to []string, message []byte) error {
	return smtp.SendMail(client.SMTPServerURL, client.Auth, client.Username, to, message)
}

func New(host, port, username, password, serverURL string) *SMTP {
	return &SMTP{
		Auth:          smtp.PlainAuth("", username, password, host),
		MIMEHeaders:   "MIME-version: 1.0;\n" + "Content-Type: text/html; charset=\"UTF-8\";\n\n",
		SMTPServerURL: host + ":" + port,
		Username:      username,
		Password:      password,
		ServerURL:     serverURL,
	}
}
