package delete

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObject"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
)

type Delete struct {
	repository.Repository
	models.Hashing
}

func (delete *Delete) Run(id *sharedVO.Id) {
	// user := delete.Repository.Search(repository.Filter{Id: id})

	// services.IsPasswordInvalid(delete.Hashing, user.Password.Value, password.Value)

	delete.Repository.Delete(id)
}

func NewDelete(repository repository.Repository, hashing models.Hashing) *Delete {
	return &Delete{
		repository,
		hashing,
	}
}
