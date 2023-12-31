package delete

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
)

type CommandHandler struct {
	Delete Delete
}

func (commandHandler *CommandHandler) Handle(command Command) {
	id := sharedVO.NewId(command.Id)

	commandHandler.Delete.Run(id)
}

func NewCommandHandler(delete Delete) *CommandHandler {
	return &CommandHandler{
		delete,
	}
}
