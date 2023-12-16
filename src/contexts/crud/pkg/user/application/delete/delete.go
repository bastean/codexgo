package delete

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type Delete struct {
	Repository repository.Repository
}

func (delete *Delete) Run(email *sharedVO.Email, password *userVO.Password) {
	delete.Repository.Delete(email)
}

func NewDelete(repository repository.Repository) *Delete {
	return &Delete{
		repository,
	}
}
