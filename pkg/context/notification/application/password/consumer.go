package password

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type Consumer struct {
	cases.Password
}

func (c *Consumer) On(event *messages.Message) error {
	account, ok := event.Attributes.(*events.UserResetQueuedAttributes)

	if !ok {
		return errors.EventAssertion("On")
	}

	err := c.Password.Run(account)

	if err != nil {
		return errors.BubbleUp(err, "On")
	}

	return nil
}
