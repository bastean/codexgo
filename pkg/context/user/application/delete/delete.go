package delete

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type Delete struct {
	model.Repository
	model.Hashing
}

func (delete *Delete) Run(id *sharedValueObject.Id) {
	// user := delete.Repository.Search(repository.Filter{Id: id})

	// service.IsPasswordInvalid(delete.Hashing, user.Password.Value, password.Value)

	delete.Repository.Delete(id)
}

func NewDelete(repository model.Repository, hashing model.Hashing) *Delete {
	return &Delete{
		repository,
		hashing,
	}
}
