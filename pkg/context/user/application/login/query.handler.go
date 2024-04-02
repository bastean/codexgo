package login

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type QueryHandler struct {
	*Login
}

func (handler *QueryHandler) Handle(query *Query) *Response {
	email := valueObject.NewEmail(query.Email)
	password := valueObject.NewPassword(query.Password)

	user := handler.Login.Run(email, password)

	response := Response(*user.ToPrimitives())

	return &response
}

func NewQueryHandler(login *Login) *QueryHandler {
	return &QueryHandler{
		Login: login,
	}
}
