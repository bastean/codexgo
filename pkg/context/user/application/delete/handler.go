package delete

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
	Action:  "delete",
	Status:  messages.Status.Queued,
}))

type CommandAttributes struct {
	ID, Password string
}

type CommandMeta struct{}

type Handler struct {
	cases.Delete
}

func (h *Handler) Handle(command *messages.Message) error {
	attributes, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	id, errID := values.New[*user.ID](attributes.ID)

	plain, errPlain := values.New[*user.PlainPassword](attributes.Password)

	err := errors.Join(errID, errPlain)

	if err != nil {
		return errors.BubbleUp(err)
	}

	err = h.Delete.Run(id, plain)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
