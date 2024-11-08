package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

const CommandType commands.Type = "user.command.creating.user"

type Command struct {
	Id, Email, Username, Password string
}

func (*Command) Type() commands.Type {
	return CommandType
}

type Handler struct {
	cases.Create
	events.Bus
}

func (handler *Handler) SubscribedTo() commands.Type {
	return CommandType
}

func (handler *Handler) Handle(cmd commands.Command) error {
	data, ok := cmd.(*Command)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	account, err := user.New(&user.Primitive{
		Id:       data.Id,
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
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
