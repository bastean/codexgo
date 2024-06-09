package login

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
)

type Login struct {
	model.Repository
	model.Hashing
}

func (login *Login) Run(input *Input) (*aggregate.User, error) {
	user, err := login.Repository.Search(&model.RepositorySearchCriteria{
		Email: input.Email,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(login.Hashing, user.Password.Value(), input.Password.Value())

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return user, nil
}
