package stransport

import (
	"net/smtp"
)

type SMTP struct {
	smtp.Auth
	MIMEHeaders                       string
	SmtpServerUrl, Username, Password string
	ServerUrl                         string
}

func NewSMTP(host, port, username, password, serverUrl string) *SMTP {
	return &SMTP{
		Auth:          smtp.PlainAuth("", username, password, host),
		MIMEHeaders:   "MIME-version: 1.0;\n" + "Content-Type: text/html; charset=\"UTF-8\";\n\n",
		SmtpServerUrl: host + ":" + port,
		Username:      username,
		Password:      password,
		ServerUrl:     serverUrl,
	}
}
