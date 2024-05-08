package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
)

type CommandHandler struct {
	model.UseCase[*Command, *types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	handler.UseCase.Run(command)

	return nil
}
