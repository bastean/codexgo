package delete

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type CommandHandler struct {
	*Delete
}

func (commandHandler *CommandHandler) Handle(command *Command) {
	id := valueObject.NewId(command.Id)

	commandHandler.Delete.Run(id)
}

func NewCommandHandler(delete *Delete) *CommandHandler {
	return &CommandHandler{
		Delete: delete,
	}
}
