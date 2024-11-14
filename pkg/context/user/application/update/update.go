package update

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/hashes"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Case struct {
	repository.Repository
	hashes.Hashing
}

func (c *Case) Run(account *user.User, updated *user.Password) error {
	found, err := c.Repository.Search(&repository.SearchCriteria{
		ID: account.ID,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = hashes.IsPasswordInvalid(c.Hashing, found.Password.Value, account.Password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if updated != nil {
		account.Password = updated
	}

	account.Verified = found.Verified

	err = c.Repository.Update(account)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
