package delete

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
)

type Delete struct {
	model.Repository
	model.Hashing
}

func (delete *Delete) Run(id *user.Id, password *user.Password) error {
	found, err := delete.Repository.Search(&model.RepositorySearchCriteria{
		Id: id,
	})

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = service.IsPasswordInvalid(delete.Hashing, found.Password.Value, password.Value)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	err = delete.Repository.Delete(found.Id)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}
