package forgot

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
	roles.EventBus
}

func (c *Case) Run(attributes *CommandAttributes) error {
	resetToken, errResetToken := values.New[*user.ID](attributes.ResetToken)

	email, errEmail := values.New[*user.Email](attributes.Email)

	err := errors.Join(errResetToken, errEmail)

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate, err := c.Repository.Search(&user.Criteria{
		Email: email,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	if aggregate.HasResetToken() {
		return nil
	}

	aggregate.ResetToken = resetToken

	err = c.Repository.Update(aggregate)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = c.EventBus.Publish(messages.New(
		user.ResetQueuedKey,
		&user.ResetQueuedAttributes{
			ResetToken: aggregate.ResetToken.Value(),
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
