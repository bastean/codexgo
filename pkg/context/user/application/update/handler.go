package update

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
	Action:  "update",
	Status:  messages.Status.Queued,
}))

type CommandAttributes struct {
	ID, Email, Username, Password, UpdatedPassword string
}

type CommandMeta struct{}

type Handler struct {
	cases.Update
}

func (h *Handler) Handle(command *messages.Message) error {
	attributes, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	aggregate, err := user.FromRaw(&user.Primitive{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.Password,
	})

	if err != nil {
		return errors.BubbleUp(err)
	}

	var updated *user.PlainPassword

	if attributes.UpdatedPassword != "" {
		updated, err = values.New[*user.PlainPassword](attributes.UpdatedPassword)

		if err != nil {
			return errors.BubbleUp(err)
		}
	}

	err = h.Update.Run(aggregate, updated)

	if err != nil {
		return errors.BubbleUp(err)
	}

	return nil
}
