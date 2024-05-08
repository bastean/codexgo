package delete

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
	id, err := valueObject.NewId(command.Id)

	if err != nil {
		return errs.BubbleUp("Handle", err)
	}

	_, err = handler.UseCase.Run(id)

	if err != nil {
		return errs.BubbleUp("Handle", err)
	}

	return nil
}
