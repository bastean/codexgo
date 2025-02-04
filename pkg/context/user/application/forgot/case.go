package forgot

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
}

func (c *Case) Run(reset *user.ID, email *user.Email) (*user.User, error) {
	aggregate, err := c.Repository.Search(&user.Criteria{
		Email: email,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if aggregate.HasReset() {
		return aggregate, nil
	}

	aggregate.Reset = reset

	err = c.Repository.Update(aggregate)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	aggregate.Record(messages.New(
		events.UserResetQueuedKey,
		&events.UserResetQueuedAttributes{
			Reset:    aggregate.Reset.Value,
			ID:       aggregate.ID.Value,
			Email:    aggregate.Email.Value,
			Username: aggregate.Username.Value,
		},
		new(events.UserResetQueuedMeta),
	))

	return aggregate, nil
}
