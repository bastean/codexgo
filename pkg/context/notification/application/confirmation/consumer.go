package confirmation

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type Consumer struct {
	cases.Confirmation
}

func (c *Consumer) On(event *messages.Message) error {
	account, ok := event.Attributes.(*events.UserCreatedSucceededAttributes)

	if !ok {
		return errors.EventAssertion("On")
	}

	err := c.Confirmation.Run(account)

	if err != nil {
		return errors.BubbleUp(err, "On")
	}

	return nil
}
