package user

import (
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type ReadQuery = read.Query

type ReadResponse = read.Response

func NewRead(repository model.Repository) *read.Handler {
	usecase := &read.Read{
		Repository: repository,
	}

	return &read.Handler{
		UseCase: usecase,
	}
}
