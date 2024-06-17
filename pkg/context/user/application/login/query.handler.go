package login

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Input struct {
	Email, Password models.ValueObject[string]
}

type Handler struct {
	models.UseCase[*Input, *aggregate.User]
}

func (handler *Handler) Handle(query *Query) (*Response, error) {
	email, errEmail := valueobj.NewEmail(query.Email)
	password, errPassword := valueobj.NewPassword(query.Password)

	err := errors.Join(errEmail, errPassword)

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	user, err := handler.UseCase.Run(&Input{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Handle")
	}

	response := Response(*user.ToPrimitives())

	return &response, nil
}
