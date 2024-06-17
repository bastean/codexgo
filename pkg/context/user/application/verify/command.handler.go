package verify

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Handler struct {
	models.UseCase[models.ValueObject[string], types.Empty]
}

func (handler *Handler) Handle(command *Command) error {
	idVO, err := valueobj.NewId(command.Id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(idVO)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
