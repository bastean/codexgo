package forgot

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

var CommandKey = messages.NewKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Command,
	Entity:  "user",
	Command: "forgot",
	Status:  messages.Status.Queued,
})

type CommandAttributes struct {
	Reset, Email string
}

type CommandMeta struct{}

type Handler struct {
	cases.Forgot
	roles.EventBus
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

	email, err := user.NewEmail(attributes.Email)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	aggregate, err := h.Forgot.Run(reset, email)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	for _, event := range aggregate.Pull() {
		err = h.EventBus.Publish(event)

		if err != nil {
			return errors.BubbleUp(err, "Handle")
		}
	}

	return nil
}
