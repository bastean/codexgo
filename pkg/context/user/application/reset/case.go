package reset

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
	roles.Hasher
}

func (c *Case) Run(attributes *CommandAttributes) error {
	resetToken, errResetToken := values.New[*user.ID](attributes.ResetToken)
	id, errID := values.New[*user.ID](attributes.ID)
	plainPassword, errPlainPassword := values.New[*user.PlainPassword](attributes.Password)

	if err := errors.Join(errResetToken, errID, errPlainPassword); err != nil {
		return errors.BubbleUp(err)
	}

	aggregate, err := c.Repository.Search(&user.Criteria{
		ID: id,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = aggregate.ValidateResetToken(resetToken)

	if err != nil {
		return errors.BubbleUp(err)
	}

	hashed, err := c.Hasher.Hash(plainPassword.Value())

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate.Password, err = values.Replace(aggregate.Password, hashed)

	if err != nil {
		return errors.BubbleUp(err)
	}

	aggregate.ResetToken = nil

	err = c.Repository.Update(aggregate)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
