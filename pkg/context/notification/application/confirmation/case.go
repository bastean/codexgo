package confirmation

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/transfer"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
)

type Case struct {
	transfer.Transfer[*user.CreatedSucceededAttributes]
}

func (c *Case) Run(event *user.CreatedSucceededAttributes) error {
	err := c.Transfer.Submit(event)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
