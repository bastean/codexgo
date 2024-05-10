package verify

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Verify struct {
	model.Repository
}

func (verify *Verify) Run(id smodel.ValueObject[string]) (*stype.Empty, error) {
	userRegistered, err := verify.Repository.Search(model.RepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return nil, serror.BubbleUp(err, "Run")
	}

	if userRegistered.Verified.Value() {
		return nil, nil
	}

	userRegistered.Verified, err = valueobj.NewVerified(true)

	if err != nil {
		return nil, serror.BubbleUp(err, "Run")
	}

	userRegistered.Password = nil

	err = verify.Repository.Update(userRegistered)

	if err != nil {
		return nil, serror.BubbleUp(err, "Run")
	}

	return nil, nil
}
