package read

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
)

type Read struct {
	repository.User
}

func (read *Read) Run(id *user.Id) (*user.User, error) {
	found, err := read.User.Search(&repository.UserSearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return found, nil
}
