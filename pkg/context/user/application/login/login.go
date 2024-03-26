package login

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/service"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Login struct {
	model.Repository
	model.Hashing
}

func (login *Login) Run(email *sharedValueObject.Email, password *valueObject.Password) *aggregate.User {
	user := login.Repository.Search(model.RepositorySearchFilter{Email: email})

	service.IsPasswordInvalid(login.Hashing, user.Password.Value, password.Value)

	return user
}

func NewLogin(repository model.Repository, hashing model.Hashing) *Login {
	return &Login{
		Repository: repository,
		Hashing:    hashing,
	}
}
