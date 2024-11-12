package delete

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/service"
)

type Case struct {
	repository.Repository
	hashing.Hashing
}

func (c *Case) Run(id *user.ID, password *user.Password) error {
	found, err := c.Repository.Search(&repository.SearchCriteria{
		ID: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(c.Hashing, found.Password.Value, password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = c.Repository.Delete(found.ID)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
