package create

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
	Command: "create",
	Status:  messages.Status.Queued,
})

type CommandAttributes struct {
	Verify, ID, Email, Username, Password string
}

type CommandMeta struct{}

type Handler struct {
	cases.Create
	roles.EventBus
}

func (h *Handler) Handle(command *messages.Message) error {
	attributes, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	aggregate, err := user.New(&user.Primitive{
		Verify:   attributes.Verify,
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.Password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = h.Create.Run(aggregate)

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
