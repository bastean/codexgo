package login

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObject"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/models"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"github.com/bastean/codexgo/context/pkg/user/domain/services"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObject"
)

type Login struct {
	repository.Repository
	models.Hashing
}

func (login *Login) Run(email *sharedVO.Email, password *userVO.Password) *aggregate.User {
	user := login.Repository.Search(repository.Filter{Email: email})

	services.IsPasswordInvalid(login.Hashing, user.Password.Value, password.Value)

	return user
}

func NewLogin(repository repository.Repository, hashing models.Hashing) *Login {
	return &Login{
		repository,
		hashing,
	}
}
