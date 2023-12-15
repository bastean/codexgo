package register

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
)

type CommandHandler struct {
	Register Register
}

func (commandHandler *CommandHandler) Handle(command Command) error {
	user, err := aggregate.Create(command.Id, command.Email, command.Username, command.Password)

	if err != nil {
		return err
	}

	commandHandler.Register.Run(user)

	return nil
}

func NewCommandHandler(register Register) *CommandHandler {
	return &CommandHandler{
		register,
	}
}
