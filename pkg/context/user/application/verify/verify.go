package verify

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Verify struct {
	model.Repository
}

func (verify *Verify) Run(id models.ValueObject[string]) (types.Empty, error) {
	user, err := verify.Repository.Search(&model.RepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	if user.Verified.Value() {
		return nil, nil
	}

	err = verify.Repository.Verify(id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return nil, nil
}
