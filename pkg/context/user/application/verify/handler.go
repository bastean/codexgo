package verify

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
	Command: "verify",
	Status:  messages.Status.Queued,
})

type CommandAttributes struct {
	Verify, ID string
}

type CommandMeta struct{}

type Handler struct {
	cases.Verify
}

func (h *Handler) Handle(command *messages.Message) error {
	attributes, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	verify, err := user.NewID(attributes.Verify)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	id, err := user.NewID(attributes.ID)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = h.Verify.Run(verify, id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
