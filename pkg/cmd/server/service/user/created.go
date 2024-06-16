package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/transports"
	"github.com/bastean/codexgo/pkg/context/user/application/created"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/communication/mail"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/communication/terminal"
)

var Created = new(created.Created)

func NewMailConfirmation(smtp *transports.SMTP) models.Transport {
	return &mail.Confirmation{
		SMTP: smtp,
	}
}

func NewTerminalConfirmation(logger models.Logger, serverURL string) models.Transport {
	return &terminal.Confirmation{
		Logger:    logger,
		ServerURL: serverURL,
	}
}

func InitCreated(transport models.Transport) error {
	Created.Transport = transport
	return nil
}
