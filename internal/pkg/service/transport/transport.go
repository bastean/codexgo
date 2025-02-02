package transport

import (
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports/smtp"
)

var Service = &struct {
	SMTP, Terminal string
}{
	SMTP:     log.Service("SMTP"),
	Terminal: log.Service("Terminal"),
}

var (
	SMTP *smtp.SMTP
)

func Up() error {
	switch {
	case env.HasSMTP():
		log.EstablishingConnectionWith(Service.SMTP)

		SMTP = smtp.Open(
			&smtp.Auth{
				Host:     env.SMTPHost,
				Port:     env.SMTPPort,
				Username: env.SMTPUsername,
				Password: env.SMTPPassword,
			},
		)

		log.ConnectionEstablishedWith(Service.SMTP)
	default:
		log.Starting(Service.Terminal)
		log.Started(Service.Terminal)
	}

	return nil
}
