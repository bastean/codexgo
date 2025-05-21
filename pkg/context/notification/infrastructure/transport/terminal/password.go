package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type Password struct {
	roles.Logger
	AppServerURL string
}

func (p *Password) Submit(recipient *recipient.Recipient) error {
	link := fmt.Sprintf("Hi %s, please reset your password through this link: %s/reset?token=%s&id=%s", recipient.Username.Value(), p.AppServerURL, recipient.ResetToken.Value(), recipient.ID.Value())

	p.Logger.Info(link)

	return nil
}
