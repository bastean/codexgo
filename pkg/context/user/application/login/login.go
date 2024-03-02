package login

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
	userVO "github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Login struct {
	repository.Repository
	model.Hashing
}

func (login *Login) Run(email *sharedVO.Email, password *userVO.Password) *aggregate.User {
	user := login.Repository.Search(repository.Filter{Email: email})

	service.IsPasswordInvalid(login.Hashing, user.Password.Value, password.Value)

	return user
}

func NewLogin(repository repository.Repository, hashing model.Hashing) *Login {
	return &Login{
		repository,
		hashing,
	}
}
