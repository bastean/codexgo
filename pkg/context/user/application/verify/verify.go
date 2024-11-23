package verify

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Case struct {
	repository.Repository
}

func (c *Case) Run(id *user.ID) error {
	aggregate, err := c.Repository.Search(&repository.SearchCriteria{
		ID: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if aggregate.Verified.Value {
		return nil
	}

	err = c.Repository.Verify(id)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
