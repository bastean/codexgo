package verify

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Verify struct {
	model.Repository
}

func (verify *Verify) Run(id *user.Id) error {
	found, err := verify.Repository.Search(&model.RepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	if found.Verified.Value {
		return nil
	}

	err = verify.Repository.Verify(id)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
