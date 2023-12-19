package login

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type Login struct {
	Repository repository.Repository
}

func (login *Login) Run(email *sharedVO.Email, password *userVO.Password) *aggregate.User {
	user, err := login.Repository.Search(email)

	if err != nil {
		panic(err.Error())
	}

	return user
}

func NewLogin(repository repository.Repository) *Login {
	return &Login{
		repository,
	}
}
