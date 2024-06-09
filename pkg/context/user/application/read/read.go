package read

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Read struct {
	model.Repository
}

func (read *Read) Run(id models.ValueObject[string]) (*aggregate.User, error) {
	user, err := read.Repository.Search(&model.RepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return user, nil
}
