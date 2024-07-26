package smtp

import (
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports/smtp"
)

type SMTP = *smtp.SMTP

func Open(host, port, username, password, serverURL string) *smtp.SMTP {
	return smtp.Open(host, port, username, password, serverURL)
}
