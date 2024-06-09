package transports

import (
	"net/smtp"
)

type SMTP struct {
	smtp.Auth
	MIMEHeaders, SMTPServerURL, Username, Password, ServerURL string
}

func NewSMTP(host, port, username, password, serverURL string) *SMTP {
	return &SMTP{
		Auth:          smtp.PlainAuth("", username, password, host),
		MIMEHeaders:   "MIME-version: 1.0;\n" + "Content-Type: text/html; charset=\"UTF-8\";\n\n",
		SMTPServerURL: host + ":" + port,
		Username:      username,
		Password:      password,
		ServerURL:     serverURL,
	}
}
