package delete

import (
	sharedVO "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
)

type CommandHandler struct {
	*Delete
}

func (commandHandler *CommandHandler) Handle(command *Command) {
	id := sharedVO.NewId(command.Id)

	commandHandler.Delete.Run(id)
}

func NewCommandHandler(delete *Delete) *CommandHandler {
	return &CommandHandler{
		delete,
	}
}
