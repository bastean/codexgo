package read

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

type Handler struct {
	cases.Read
}

func (handler *Handler) Handle(query *Query) (*Response, error) {
	id, err := user.NewId(query.Id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	found, err := handler.Read.Run(id)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*found.ToPrimitive())

	return &response, nil
}
