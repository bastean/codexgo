package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Case struct {
	repository.Repository
}

func (c *Case) Run(account *user.User) error {
	err := c.Repository.Create(account)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
