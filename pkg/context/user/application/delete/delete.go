package delete

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Delete struct {
	model.Repository
	model.Hashing
}

func (delete *Delete) Run(id smodel.ValueObject[string]) (*stype.Empty, error) {
	// TODO!: user := delete.Repository.Search(repository.Filter{Id: id})

	// TODO!: service.IsPasswordInvalid(delete.Hashing, user.Password.Value, password.Value)

	err := delete.Repository.Delete(id)

	if err != nil {
		return nil, serror.BubbleUp(err, "Run")
	}

	return nil, nil
}
