package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type Confirmation struct {
	roles.Logger
	AppServerURL string
}

func (c *Confirmation) Submit(attributes *events.UserCreatedSucceededAttributes) error {
	link := fmt.Sprintf("Hi %s, please confirm your account through this link: %s/v4/account/verify?token=%s&id=%s", attributes.Username, c.AppServerURL, attributes.Verify, attributes.ID)

	c.Logger.Info(link)

	return nil
}
