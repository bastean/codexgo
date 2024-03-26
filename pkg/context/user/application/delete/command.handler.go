package delete

import (
	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
)

type CommandHandler struct {
	*Delete
}

func (commandHandler *CommandHandler) Handle(command *Command) {
	id := sharedValueObject.NewId(command.Id)

	commandHandler.Delete.Run(id)
}

func NewCommandHandler(delete *Delete) *CommandHandler {
	return &CommandHandler{
		Delete: delete,
	}
}
