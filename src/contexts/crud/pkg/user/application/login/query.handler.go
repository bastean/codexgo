package login

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type QueryHandler struct {
	Login Login
}

func (commandHandler *QueryHandler) Handle(query Query) error {
	email, err := sharedVO.NewEmail(query.Email)

	if err != nil {
		return err
	}

	password, err := userVO.NewPassword(query.Password)

	if err != nil {
		return err
	}

	commandHandler.Login.Run(email, password)

	return nil
}

func NewQueryHandler(login Login) *QueryHandler {
	return &QueryHandler{
		login,
	}
}
