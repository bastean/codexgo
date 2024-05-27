package user

import (
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type DeleteCommand = delete.Command

func NewDelete(repository model.Repository, hashing model.Hashing) *delete.CommandHandler {
	useCase := &delete.Delete{
		Repository: repository,
		Hashing:    hashing,
	}

	return &delete.CommandHandler{
		UseCase: useCase,
	}
}
