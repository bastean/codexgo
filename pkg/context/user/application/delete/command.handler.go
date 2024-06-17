package delete

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Input struct {
	Id, Password models.ValueObject[string]
}

type Handler struct {
	models.UseCase[*Input, types.Empty]
}

func (handler *Handler) Handle(command *Command) error {
	id, errId := valueobj.NewId(command.Id)
	password, errPassword := valueobj.NewPassword(command.Password)

	err := errors.Join(errId, errPassword)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(&Input{
		Id:       id,
		Password: password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
