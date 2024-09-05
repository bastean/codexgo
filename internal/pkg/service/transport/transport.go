package transport

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
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

		user.InitCreated(&user.MailConfirmation{
			SMTP:         SMTP,
			AppServerURL: env.ServerGinURL,
		}, user.QueueSendConfirmation)

		log.ConnectionEstablishedWith(Service.SMTP)
	default:
		log.Starting(Service.Terminal)

		user.InitCreated(&user.TerminalConfirmation{
			Logger:       log.Log,
			AppServerURL: env.ServerGinURL,
		}, user.QueueSendConfirmation)

		log.Started(Service.Terminal)
	}

	return nil
}
