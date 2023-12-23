package delete

import (
	sharedVO "github.com/bastean/codexgo/context/pkg/shared/domain/valueObjects"
	userVO "github.com/bastean/codexgo/context/pkg/user/domain/valueObjects"
)

type CommandHandler struct {
	Delete Delete
}

func (commandHandler *CommandHandler) Handle(command Command) {
	id := sharedVO.NewId(command.Id)
	password := userVO.NewPassword(command.Password)

	commandHandler.Delete.Run(id, password)
}

func NewCommandHandler(delete Delete) *CommandHandler {
	return &CommandHandler{
		delete,
	}
}
