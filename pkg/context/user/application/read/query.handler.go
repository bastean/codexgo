package read

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Input struct {
	Id smodel.ValueObject[string]
}

type QueryHandler struct {
	smodel.UseCase[*Input, *aggregate.User]
}

func (handler *QueryHandler) Handle(query *Query) (*Response, error) {
	id, err := valueobj.NewId(query.Id)

	if err != nil {
		return nil, serror.BubbleUp(err, "Handle")
	}

	user, err := handler.UseCase.Run(&Input{
		Id: id,
	})

	if err != nil {
		return nil, serror.BubbleUp(err, "Handle")
	}

	response := Response(*user.ToPrimitives())

	return &response, nil
}
