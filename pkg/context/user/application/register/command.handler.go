package register

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type CommandHandler struct {
	*Register
	model.Broker
}

func (commandHandler *CommandHandler) Handle(command *Command) {
	user := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	commandHandler.Register.Run(user)

	commandHandler.Broker.PublishMessages(user.PullMessages())
}

func NewCommandHandler(register *Register, broker model.Broker) *CommandHandler {
	return &CommandHandler{
		Register: register,
		Broker:   broker,
	}
}
