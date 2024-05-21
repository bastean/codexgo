package login

import (
	"errors"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Input struct {
	Email    smodel.ValueObject[string]
	Password smodel.ValueObject[string]
}

type QueryHandler struct {
	smodel.UseCase[*Input, *aggregate.User]
}

func (handler *QueryHandler) Handle(query *Query) (*Response, error) {
	email, errEmail := valueobj.NewEmail(query.Email)
	password, errPassword := valueobj.NewPassword(query.Password)

	err := errors.Join(errEmail, errPassword)

	if err != nil {
		return nil, serror.BubbleUp(err, "Handle")
	}

	user, err := handler.UseCase.Run(&Input{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, serror.BubbleUp(err, "Handle")
	}

	response := Response(*user.ToPrimitives())

	return &response, nil
}
