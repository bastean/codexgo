package reset

import (
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
	Command: "reset",
	Status:  messages.Status.Queued,
})

type CommandAttributes struct {
	Reset, ID, Password string
}

type CommandMeta struct{}

type Handler struct {
	cases.Reset
}

func (h *Handler) Handle(command *messages.Message) error {
	attributes, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	reset, err := user.NewID(attributes.Reset)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	id, err := user.NewID(attributes.ID)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	password, err := user.NewPlainPassword(attributes.Password)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = h.Reset.Run(reset, id, password)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
