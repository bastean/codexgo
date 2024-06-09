package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
)

type CommandHandler struct {
	models.UseCase[*Command, types.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	_, err := handler.UseCase.Run(command)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
