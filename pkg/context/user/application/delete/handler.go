package delete

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

var CommandKey = messages.NewKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Command,
	Entity:  "user",
	Command: "deleting",
	Status:  messages.Status.Queued,
})

type CommandAttributes struct {
	Id, Password string
}

type CommandMeta struct{}

type Handler struct {
	cases.Delete
}

func (handler *Handler) Handle(command *commands.Command) error {
	attributes, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	id, errId := user.NewId(attributes.Id)

	password, errPassword := user.NewPassword(attributes.Password)

	err := errors.Join(errId, errPassword)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Delete.Run(id, password)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
