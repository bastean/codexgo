package terminal

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type Password struct {
	roles.Logger
	AppServerURL string
}

func (p *Password) Submit(attributes *events.UserResetQueuedAttributes) error {
	link := fmt.Sprintf("Hi %s, please reset your password through this link: %s/reset?token=%s&id=%s", attributes.Username, p.AppServerURL, attributes.Reset, attributes.ID)

	p.Logger.Info(link)

	return nil
}
