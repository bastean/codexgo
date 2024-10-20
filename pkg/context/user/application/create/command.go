package create

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

const CommandType command.Type = "user.command.creating.user"

type Command struct {
	Id, Email, Username, Password string
}

func (*Command) Type() command.Type {
	return CommandType
}

type Handler struct {
	cases.Create
	messages.Broker
}

func (handler *Handler) SubscribedTo() command.Type {
	return CommandType
}

func (handler *Handler) Handle(cmd command.Command) error {
	data, ok := cmd.(*Command)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	aggregate, err := user.New(&user.Primitive{
		Id:       data.Id,
		Email:    data.Email,
		Username: data.Username,
		Password: data.Password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Create.Run(aggregate)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	handler.Broker.PublishMessages(aggregate.Pull())

	return nil
}
