package reset

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/role"
)

type Case struct {
	role.Repository
	roles.Hasher
}

func (c *Case) Run(reset, id *user.ID, password *user.PlainPassword) error {
	aggregate, err := c.Repository.Search(&user.Criteria{
		ID: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = aggregate.ValidateReset(reset)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	hashed, err := c.Hasher.Hash(password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	aggregate.CipherPassword, err = user.NewCipherPassword(hashed)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	aggregate.Reset = nil

	err = c.Repository.Update(aggregate)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
