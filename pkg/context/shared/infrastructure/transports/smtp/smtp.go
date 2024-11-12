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

func (s *SMTP) Headers(to, subject string) string {
	header := s.SetHeader("From", s.Username)

	header += s.SetHeader("To", to)

	header += s.SetHeader("Subject", subject)

	header += s.SetHeader("MIME-version", "1.0")

	header += s.SetHeader("Content-Type", "text/html; charset=\"UTF-8\"")

	header += "\n"

	return header
}

func (s *SMTP) SendMail(to []string, message []byte) error {
	return smtp.SendMail(
		s.ServerURL,
		s.Auth,
		s.Username,
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
