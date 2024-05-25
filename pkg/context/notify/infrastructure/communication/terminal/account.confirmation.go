package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
)

type AccountConfirmation struct {
	smodel.Logger
	ServerURL string
}

func (client *AccountConfirmation) Submit(data any) error {
	account := data.(*send.CreatedSucceededEventAttributes)

	confirmationLink := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/verify/%s", account.Username, client.ServerURL, account.Id)

	client.Logger.Info(confirmationLink)

	return nil
}
