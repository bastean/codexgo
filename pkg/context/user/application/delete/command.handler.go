package delete

import (
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
)

type CommandHandler struct {
	*Delete
}

func (handler *CommandHandler) Handle(command *Command) {
	id := valueObject.NewId(command.Id)

	handler.Delete.Run(id)
}

func NewCommandHandler(delete *Delete) *CommandHandler {
	return &CommandHandler{
		Delete: delete,
	}
}
