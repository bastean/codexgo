package delete

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"github.com/bastean/codexgo/context/pkg/user/domain/services"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type Delete struct {
	Repository repository.Repository
	Hashing    models.Hashing
}

func (delete *Delete) Run(id *sharedVO.Id, password *userVO.Password) {
	user := delete.Repository.Search(repository.Filter{Id: id})

	services.IsPasswordInvalid(delete.Hashing, user.Password.Value, password.Value)

	delete.Repository.Delete(id)
}

func NewDelete(repository repository.Repository, hashing models.Hashing) *Delete {
	return &Delete{
		repository,
		hashing,
	}
}
