package password

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type Case struct {
	role.Transfer
}

func (c *Case) Run(attributes *EventAttributes) error {
	aggregate, err := recipient.New(&recipient.Required{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate.ResetToken, err = values.New[*recipient.ID](attributes.ResetToken)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = c.Transfer.Submit(aggregate)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
