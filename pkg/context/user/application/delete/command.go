package delete

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

const CommandType command.Type = "user.command.deleting.user"

type Command struct {
	Id, Password string
}

func (*Command) Type() command.Type {
	return CommandType
}

type Handler struct {
	cases.Delete
}

func (handler *Handler) SubscribedTo() command.Type {
	return CommandType
}

func (handler *Handler) Handle(cmd command.Command) error {
	data, ok := cmd.(*Command)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	id, errId := user.NewId(data.Id)

	password, errPassword := user.NewPassword(data.Password)

	err := errors.Join(errId, errPassword)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Delete.Run(id, password)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
