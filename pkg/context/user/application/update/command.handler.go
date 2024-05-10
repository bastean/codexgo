package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
)

type CommandHandler struct {
	smodel.UseCase[*Command, *stype.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	handler.UseCase.Run(command)
	return nil
}
