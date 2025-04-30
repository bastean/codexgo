package reset

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

var CommandKey, _ = values.New[*messages.Key](messages.ParseKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Command,
	Entity:  "user",
	Action:  "reset",
	Status:  messages.Status.Queued,
}))

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

	reset, err := values.New[*user.ID](attributes.Reset)

	if err != nil {
		return errors.BubbleUp(err)
	}

	id, err := values.New[*user.ID](attributes.ID)

	if err != nil {
		return errors.BubbleUp(err)
	}

	password, err := values.New[*user.PlainPassword](attributes.Password)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = h.Reset.Run(reset, id, password)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
