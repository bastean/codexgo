package password

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

type Case struct {
	role.Transfer[*events.UserResetQueuedAttributes]
}

func (c *Case) Run(event *events.UserResetQueuedAttributes) error {
	err := c.Transfer.Submit(event)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
