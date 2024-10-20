package update

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

const CommandType command.Type = "user.command.updating.user"

type Command struct {
	Id, Email, Username, Password, UpdatedPassword string
}

func (*Command) Type() command.Type {
	return CommandType
}

type Handler struct {
	cases.Update
}

func (handler *Handler) SubscribedTo() command.Type {
	return CommandType
}

func (handler *Handler) Handle(cmd command.Command) error {
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
