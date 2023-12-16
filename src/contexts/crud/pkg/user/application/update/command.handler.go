package update

import (
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
)

type CommandHandler struct {
	Update Update
}

func (commandHandler *CommandHandler) Handle(command Command) error {
	user, err := aggregate.Create(command.Id, command.Email, command.Username, command.CurrentPassword)

	if err != nil {
		return err
	}

	commandHandler.Update.Run(user)

	return nil
}

func NewCommandHandler(update Update) *CommandHandler {
	return &CommandHandler{
		update,
	}
}
