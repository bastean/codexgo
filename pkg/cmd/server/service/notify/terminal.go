package notify

import (
	"github.com/bastean/codexgo/pkg/context/notify/domain/model"
	"github.com/bastean/codexgo/pkg/context/notify/infrastructure/communication/terminal"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
)

func NewTerminalAccountConfirmation(serverURL string, logger models.Logger) model.Transport {
	return &terminal.AccountConfirmation{
		ServerURL: serverURL,
		Logger:    logger,
	}
}
