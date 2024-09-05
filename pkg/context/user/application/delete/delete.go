package delete

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/hashing"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/service"
)

type Delete struct {
	repository.Repository
	hashing.Hashing
}

func (use *Delete) Run(id *user.Id, password *user.Password) error {
	found, err := use.Repository.Search(&repository.SearchCriteria{
		Id: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(use.Hashing, found.Password.Value, password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = use.Repository.Delete(found.Id)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
