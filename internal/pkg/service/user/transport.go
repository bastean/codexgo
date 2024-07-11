package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports/smtp"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/communication/mail"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/communication/terminal"
)

func MailConfirmation(smtp *smtp.SMTP) models.Transport {
	return &mail.Confirmation{
		SMTP: smtp,
	}
}

func TerminalConfirmation(logger models.Logger, serverURL string) models.Transport {
	return &terminal.Confirmation{
		Logger:    logger,
		ServerURL: serverURL,
	}
}
