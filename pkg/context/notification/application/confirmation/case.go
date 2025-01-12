package confirmation

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/transfer"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

type Case struct {
	transfer.Transfer[*events.UserCreatedSucceededAttributes]
}

func (c *Case) Run(event *events.UserCreatedSucceededAttributes) error {
	err := c.Transfer.Submit(event)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
