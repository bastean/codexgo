package delete

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
)

type Delete struct {
	model.Repository
	model.Hashing
}

func (delete *Delete) Run(input *Input) (types.Empty, error) {
	user, err := delete.Repository.Search(&model.RepositorySearchCriteria{
		Id: input.Id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(delete.Hashing, user.Password.Value(), input.Password.Value())

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	err = delete.Repository.Delete(user.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}
