package verify

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Verify struct {
	model.Repository
}

func (verify *Verify) Run(id sharedModel.ValueObject[string]) (*types.Empty, error) {
	userRegistered, err := verify.Repository.Search(model.RepositorySearchCriteria{Id: id})

	if err != nil {
		return nil, errs.BubbleUp("Run", err)
	}

	if userRegistered.Verified.Value() {
		return nil, nil
	}

	userRegistered.Verified, err = valueObject.NewVerified(true)

	if err != nil {
		return nil, errs.BubbleUp("Run", err)
	}

	userRegistered.Password = nil

	err = verify.Repository.Update(userRegistered)

	if err != nil {
		return nil, errs.BubbleUp("Run", err)
	}

	return nil, nil
}
