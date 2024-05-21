package read

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Read struct {
	model.Repository
}

func (read *Read) Run(input *Input) (*aggregate.User, error) {
	user, err := read.Repository.Search(model.RepositorySearchCriteria{
		Id: input.Id,
	})

	if err != nil {
		return nil, serror.BubbleUp(err, "Run")
	}

	return user, nil
}
