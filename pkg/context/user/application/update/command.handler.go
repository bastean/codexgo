package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
)

type CommandHandler struct {
	smodel.UseCase[*Command, *stype.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	_, err := handler.UseCase.Run(command)

	if err != nil {
		return serror.BubbleUp(err, "Handle")
	}

	return nil
}
