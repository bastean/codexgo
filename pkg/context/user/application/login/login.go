package login

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

func (c *Case) Run(email *user.Email, password *user.Password) (*user.User, error) {
	found, err := c.Repository.Search(&repository.SearchCriteria{
		Email: email,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(c.Hashing, found.Password.Value, password.Value)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return found, nil
}
