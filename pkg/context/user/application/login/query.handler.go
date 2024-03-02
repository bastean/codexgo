package login

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	userVO "github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type QueryHandler struct {
	Login
}

func (queryHandler *QueryHandler) Handle(query Query) *Response {
	email := sharedVO.NewEmail(query.Email)
	password := userVO.NewPassword(query.Password)

	user := queryHandler.Login.Run(email, password)

	response := Response(*user.ToPrimitives())

	return &response
}

func NewQueryHandler(login Login) *QueryHandler {
	return &QueryHandler{
		login,
	}
}
