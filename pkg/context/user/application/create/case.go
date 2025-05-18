package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	roles.Hasher
	role.Repository
	roles.EventBus
}

func (c *Case) Run(attributes *CommandAttributes) error {
	_, err := values.New[*user.PlainPassword](attributes.Password)

	if err != nil {
		return errors.BubbleUp(err)
	}

	attributes.Password, err = c.Hasher.Hash(attributes.Password)

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate, err := user.New(&user.Required{
		VerifyToken: attributes.VerifyToken,
		ID:          attributes.ID,
		Email:       attributes.Email,
		Username:    attributes.Username,
		Password:    attributes.Password,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = c.Repository.Create(aggregate)

	if err != nil {
		return errors.BubbleUp(err)
	}

	for _, event := range aggregate.Pull() {
		err = errors.Join(err, c.EventBus.Publish(event))
	}

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
