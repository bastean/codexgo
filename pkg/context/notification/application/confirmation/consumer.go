package confirmation

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type EventAttributes = struct {
	VerifyToken, ID, Email, Username string
}

type EventMeta = struct{}

type Consumer struct {
	*Case
}

func (c *Consumer) On(event *messages.Message) error {
	attributes, ok := event.Attributes.(*EventAttributes)

	if !ok {
		return errors.EventAssertion()
	}

	err := c.Case.Run(attributes)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
