package user

import (
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type LoginQuery = login.Query

func NewLogin(repository model.Repository, hashing model.Hashing) *login.QueryHandler {
	useCase := &login.Login{
		Repository: repository,
		Hashing:    hashing,
	}

	return &login.QueryHandler{
		UseCase: useCase,
	}
}
