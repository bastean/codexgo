package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type Confirmation struct {
	roles.Logger
	AppServerURL string
}

func (c *Confirmation) Submit(recipient *recipient.Recipient) error {
	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/v4/account/verify?token=%s&id=%s", recipient.Username.Value(), c.AppServerURL, recipient.VerifyToken.Value(), recipient.ID.Value())

	c.Logger.Info(link)

	return nil
}
