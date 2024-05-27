package user

import (
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type UpdateCommand = update.Command

func NewUpdate(repository model.Repository, hashing model.Hashing) *update.CommandHandler {
	useCase := &update.Update{
		Repository: repository,
		Hashing:    hashing,
	}

	return &update.CommandHandler{
		UseCase: useCase,
	}
}
