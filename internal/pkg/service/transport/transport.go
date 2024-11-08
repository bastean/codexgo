package transport

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/notification"
	"github.com/bastean/codexgo/v4/internal/pkg/service/record/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/transport/smtp"
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

		notification.Init(&notification.MailConfirmation{
			SMTP:         SMTP,
			AppServerURL: env.ServerGinURL,
		})

		log.ConnectionEstablishedWith(Service.SMTP)
	default:
		log.Starting(Service.Terminal)

		notification.Init(&notification.TerminalConfirmation{
			Logger:       log.Log,
			AppServerURL: env.ServerGinURL,
		})

		log.Started(Service.Terminal)
	}

	return nil
}
