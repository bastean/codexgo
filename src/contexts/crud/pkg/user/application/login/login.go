package login

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type Login struct {
	Repository repository.Repository
}

func (login *Login) Run(email *sharedVO.Email, password *userVO.Password) {
	login.Repository.Search(email)
}

func NewLogin(repository repository.Repository) *Login {
	return &Login{
		repository,
	}
}
