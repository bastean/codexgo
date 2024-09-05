package verify

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
)

type Verify struct {
	repository.Repository
}

func (use *Verify) Run(id *user.Id) error {
	found, err := use.Repository.Search(&repository.SearchCriteria{
		Id: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if found.Verified.Value {
		return nil
	}

	err = use.Repository.Verify(id)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
