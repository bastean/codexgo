package delete

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Delete struct {
	model.Repository
	model.Hashing
}

func (delete *Delete) Run(id sharedModel.ValueObject[string]) (*types.Empty, error) {
	// TODO!: user := delete.Repository.Search(repository.Filter{Id: id})

	// TODO!: service.IsPasswordInvalid(delete.Hashing, user.Password.Value, password.Value)

	err := delete.Repository.Delete(id)

	if err != nil {
		return nil, errs.BubbleUp("Run", err)
	}

	return nil, nil
}
