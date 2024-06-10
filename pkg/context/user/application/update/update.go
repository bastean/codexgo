package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
)

type Update struct {
	model.Repository
	model.Hashing
}

func (update *Update) Run(input *Input) (types.Empty, error) {
	user, err := update.Repository.Search(&model.RepositorySearchCriteria{
		Id: input.User.Id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(update.Hashing, user.Password.Value(), input.User.Password.Value())

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if input.UpdatedPassword != nil {
		input.User.Password = input.UpdatedPassword
	}

	input.User.Verified = user.Verified

	err = update.Repository.Update(input.User)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}
