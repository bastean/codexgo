package delete

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

func (c *Case) Run(id *user.ID, plain *user.PlainPassword) error {
	aggregate, err := c.Repository.Search(&repository.SearchCriteria{
		ID: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = hashes.IsPasswordInvalid(c.Hashing, aggregate.CipherPassword.Value, plain.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = c.Repository.Delete(aggregate.ID)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
