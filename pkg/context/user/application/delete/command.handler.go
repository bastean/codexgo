package delete

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type CommandHandler struct {
	smodel.UseCase[smodel.ValueObject[string], *stype.Empty]
}

func (handler *CommandHandler) Handle(command *Command) error {
	id, err := valueobj.NewId(command.Id)

	if err != nil {
		return serror.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(id)

	if err != nil {
		return serror.BubbleUp(err, "Handle")
	}

	return nil
}
