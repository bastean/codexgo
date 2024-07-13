package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/loggers"
	"github.com/bastean/codexgo/pkg/context/shared/domain/transfers"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports/smtp"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/transport/mail"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/transport/terminal"
)

func MailConfirmation(smtp *smtp.SMTP) transfers.Transfer {
	return &mail.Confirmation{
		SMTP: smtp,
	}
}

func TerminalConfirmation(logger loggers.Logger, serverURL string) transfers.Transfer {
	return &terminal.Confirmation{
		Logger:    logger,
		ServerURL: serverURL,
	}
}
