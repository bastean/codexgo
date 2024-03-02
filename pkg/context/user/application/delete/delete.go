package delete

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
)

type Delete struct {
	repository.Repository
	model.Hashing
}

func (delete *Delete) Run(id *sharedVO.Id) {
	// user := delete.Repository.Search(repository.Filter{Id: id})

	// service.IsPasswordInvalid(delete.Hashing, user.Password.Value, password.Value)

	delete.Repository.Delete(id)
}

func NewDelete(repository repository.Repository, hashing model.Hashing) *Delete {
	return &Delete{
		repository,
		hashing,
	}
}
