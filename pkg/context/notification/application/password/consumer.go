package password

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type EventAttributes = struct {
	ResetToken, ID, Email, Username string
}

type EventMeta = struct{}

type Consumer struct {
	*Case
}

func (c *Consumer) On(event *messages.Message) error {
	account, ok := event.Attributes.(*EventAttributes)

	if !ok {
		return errors.EventAssertion()
	}

	err := c.Case.Run(account)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
