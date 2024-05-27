package notify

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/terminal"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
)

func NewTerminalAccountConfirmation(logger models.Logger, serverURL string) model.Transport {
	return &terminal.AccountConfirmation{
		Logger:    logger,
		ServerURL: serverURL,
	}
}
