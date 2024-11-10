package verify

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
	Command: "verifying",
	Status:  messages.Status.Queued,
})

type CommandAttributes struct {
	Id string
}

type CommandMeta struct{}

type Handler struct {
	cases.Verify
}

func (handler *Handler) Handle(command *commands.Command) error {
	data, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	id, err := user.NewId(data.Id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Verify.Run(id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}