package update

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/service"
)

type Case struct {
	role.Repository
	roles.Hasher
}

func (c *Case) Run(attributes *CommandAttributes) error {
	id, errID := values.New[*values.ID](attributes.ID)
	plainPassword, errPlainPassword := values.New[*user.PlainPassword](attributes.Password)

	if err := errors.Join(errID, errPlainPassword); err != nil {
		return errors.BubbleUp(err)
	}

	aggregate, err := c.Repository.Search(&user.Criteria{
		ID: id,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = c.Hasher.Compare(aggregate.Password.Value(), plainPassword.Value())

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = service.UpdateEmail(aggregate, attributes.Email)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = service.UpdateUsername(aggregate, attributes.Username)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = service.UpdatePassword(aggregate, attributes.UpdatedPassword, c.Hasher)

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

	return nil
}
