package smtp

import (
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports/smtp"
)

type SMTP = *smtp.SMTP

func New(host, port, username, password, serverURL string) *smtp.SMTP {
	return smtp.New(host, port, username, password, serverURL)
}
