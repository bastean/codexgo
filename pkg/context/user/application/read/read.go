package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Case struct {
	repository.Repository
}

func (c *Case) Run(id *user.ID) (*user.User, error) {
	found, err := c.Repository.Search(&repository.SearchCriteria{
		ID: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return found, nil
}
