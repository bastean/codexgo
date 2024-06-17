package user

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
)

type CreateCommand = create.Command

func NewCreate(repository model.Repository, broker models.Broker) *create.Handler {
	usecase := &create.Create{
		Repository: repository,
	}

	return &create.Handler{
		UseCase: usecase,
		Broker:  broker,
	}
}
