package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type CreateCommand = create.Command

func NewCreate(repository model.Repository, broker models.Broker) *create.CommandHandler {
	useCase := &create.Create{
		Repository: repository,
	}

	return &create.CommandHandler{
		UseCase: useCase,
		Broker:  broker,
	}
}
