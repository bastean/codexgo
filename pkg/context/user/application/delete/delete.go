package delete

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Delete struct {
	model.Repository
	model.Hashing
}

func (delete *Delete) Run(id *valueObject.Id) {
	// TODO: user := delete.Repository.Search(repository.Filter{Id: id})

	// TODO: service.IsPasswordInvalid(delete.Hashing, user.Password.Value, password.Value)

	delete.Repository.Delete(id)
}

func NewDelete(repository model.Repository, hashing model.Hashing) *Delete {
	return &Delete{
		Repository: repository,
		Hashing:    hashing,
	}
}
