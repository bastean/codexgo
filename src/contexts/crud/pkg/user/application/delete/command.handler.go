package delete

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type CommandHandler struct {
	Delete Delete
}

func (commandHandler *CommandHandler) Handle(command Command) error {
	email, err := sharedVO.NewEmail(command.Email)

	if err != nil {
		return err
	}

	password, err := userVO.NewPassword(command.Password)

	if err != nil {
		return err
	}

	commandHandler.Delete.Run(email, password)

	return nil
}

func NewCommandHandler(delete Delete) *CommandHandler {
	return &CommandHandler{
		delete,
	}
}
