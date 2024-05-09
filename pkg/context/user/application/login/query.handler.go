package login

import (
	"errors"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type Input struct {
	Email    sharedModel.ValueObject[string]
	Password sharedModel.ValueObject[string]
}

type QueryHandler struct {
	sharedModel.UseCase[*Input, *aggregate.User]
}

func (handler *QueryHandler) Handle(query *Query) (*Response, error) {
	email, emailErr := valueObject.NewEmail(query.Email)
	password, passwordErr := valueObject.NewPassword(query.Password)

	err := errors.Join(emailErr, passwordErr)

	if err != nil {
		return nil, errs.BubbleUp(err, "Handle")
	}

	user, err := handler.UseCase.Run(&Input{
		Email:    email,
		Password: password,
	})

	if err != nil {
		return nil, errs.BubbleUp(err, "Handle")
	}

	response := Response(*user.ToPrimitives())

	return &response, nil
}
