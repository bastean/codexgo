package delete

import (
	"errors"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Input struct {
	Id       smodel.ValueObject[string]
	Password smodel.ValueObject[string]
}

type CommandHandler struct {
	smodel.UseCase[*Input, *stype.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	id, errId := valueobj.NewId(command.Id)
	password, errPassword := valueobj.NewPassword(command.Password)

	err := errors.Join(errId, errPassword)

	if err != nil {
		return serror.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(&Input{
		Id:       id,
		Password: password,
	})

	if err != nil {
		return serror.BubbleUp(err, "Handle")
	}

	return nil
}
