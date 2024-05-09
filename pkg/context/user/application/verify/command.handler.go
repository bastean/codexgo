package verify

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type CommandHandler struct {
	model.UseCase[model.ValueObject[string], *types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	idVO, err := valueObject.NewId(command.Id)

	if err != nil {
		return errs.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(idVO)

	if err != nil {
		return errs.BubbleUp(err, "Handle")
	}

	return nil
}
