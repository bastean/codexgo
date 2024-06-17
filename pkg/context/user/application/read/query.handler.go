package read

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Handler struct {
	models.UseCase[models.ValueObject[string], *aggregate.User]
}

func (handler *Handler) Handle(query *Query) (*Response, error) {
	id, err := valueobj.NewId(query.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	user, err := handler.UseCase.Run(id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*user.ToPrimitives())

	return &response, nil
}
