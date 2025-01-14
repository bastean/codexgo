package update

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

func (c *Case) Run(aggregate *user.User, updated *user.PlainPassword) error {
	account, err := c.Repository.Search(&user.Criteria{
		ID: aggregate.ID,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = c.Hasher.Compare(account.CipherPassword.Value, aggregate.PlainPassword.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	hashed := account.CipherPassword.Value

	if updated != nil {
		hashed, err = c.Hasher.Hash(updated.Value)

		if err != nil {
			return errors.BubbleUp(err, "Run")
		}
	}

	aggregate.CipherPassword, err = user.NewCipherPassword(hashed)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	aggregate.Created = account.Created
	aggregate.Updated = account.Updated
	aggregate.Verified = account.Verified

	err = c.Repository.Update(aggregate)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
