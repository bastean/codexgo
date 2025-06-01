package forgot

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/service"
)

type Case struct {
	role.Repository
	roles.EventBus
}

func (c *Case) Run(attributes *CommandAttributes) error {
	email, err := values.New[*values.Email](attributes.Email)

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate, err := c.Repository.Search(&user.Criteria{
		Email: email,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = service.SetResetToken(aggregate, attributes.ResetToken)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = aggregate.UpdatedStamp()

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = c.Repository.Update(aggregate)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = c.EventBus.Publish(messages.New(
		user.ResetQueuedKey,
		&user.ResetQueuedAttributes{
			ResetToken: aggregate.ResetToken.ID.Value(),
			ID:         aggregate.ID.Value(),
			Email:      aggregate.Email.Value(),
			Username:   aggregate.Username.Value(),
		},
		new(user.ResetQueuedMeta),
	))

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
