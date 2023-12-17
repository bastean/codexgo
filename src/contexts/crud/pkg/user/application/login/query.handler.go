package login

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type QueryHandler struct {
	Login Login
}

func (commandHandler *QueryHandler) Handle(query Query) error {
	email := sharedVO.NewEmail(query.Email)
	password := userVO.NewPassword(query.Password)

	commandHandler.Login.Run(email, password)

	return nil
}

func NewQueryHandler(login Login) *QueryHandler {
	return &QueryHandler{
		login,
	}
}
