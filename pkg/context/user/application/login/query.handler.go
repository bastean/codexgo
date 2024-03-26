package login

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	valueObject "github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type QueryHandler struct {
	*Login
}

func (queryHandler *QueryHandler) Handle(query *Query) *Response {
	email := sharedValueObject.NewEmail(query.Email)
	password := valueObject.NewPassword(query.Password)

	user := queryHandler.Login.Run(email, password)

	response := Response(*user.ToPrimitives())

	return &response
}

func NewQueryHandler(login *Login) *QueryHandler {
	return &QueryHandler{
		Login: login,
	}
}
