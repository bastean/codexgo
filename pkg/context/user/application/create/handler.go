package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

var CommandKey = messages.NewKey(&messages.KeyComponents{
	Service: "user",
	Version: "1",
	Type:    messages.Type.Command,
	Entity:  "user",
	Command: "creating",
	Status:  messages.Status.Queued,
})

type CommandAttributes struct {
	Id, Email, Username, Password string
}

type CommandMeta struct{}

type Handler struct {
	cases.Create
	events.Bus
}

func (handler *Handler) Handle(command *commands.Command) error {
	attributes, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	account, err := user.New(&user.Primitive{
		Id:       attributes.Id,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.Password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Create.Run(account)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	for _, event := range account.Pull() {
		err = handler.Bus.Publish(event)

		if err != nil {
			return errors.BubbleUp(err, "Handle")
		}
	}

	return nil
}