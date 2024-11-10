package update

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
	Command: "updating",
	Status:  messages.Status.Queued,
})

type CommandAttributes struct {
	Id, Email, Username, Password, UpdatedPassword string
}

type CommandMeta struct{}

type Handler struct {
	cases.Update
}

func (handler *Handler) Handle(command *commands.Command) error {
	data, ok := command.Attributes.(*CommandAttributes)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	account, err := user.FromPrimitive(&user.Primitive{
		Id:       data.Id,
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	var updated *user.Password

	if data.UpdatedPassword != "" {
		updated, err = user.NewPassword(data.UpdatedPassword)

		if err != nil {
			return errors.BubbleUp(err, "Handle")
		}
	}

	err = handler.Update.Run(account, updated)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
