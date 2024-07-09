package read

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Read struct {
	model.Repository
}

func (read *Read) Run(id *user.Id) (*user.User, error) {
	found, err := read.Repository.Search(&model.RepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return found, nil
}
