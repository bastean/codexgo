package register

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
)

type CommandHandler struct {
	Register Register
}

func (commandHandler *CommandHandler) Handle(command Command) {
	user := aggregate.Create(command.Id, command.Email, command.Username, command.Password)

	commandHandler.Register.Run(user)
}

func NewCommandHandler(register Register) *CommandHandler {
	return &CommandHandler{
		register,
	}
}
