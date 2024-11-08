package verify

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
)

const CommandType commands.Type = "user.command.verifying.user"

type Command struct {
	Id string
}

func (*Command) Type() commands.Type {
	return CommandType
}

type Handler struct {
	cases.Verify
}

func (handler *Handler) SubscribedTo() commands.Type {
	return CommandType
}

func (handler *Handler) Handle(cmd commands.Command) error {
	data, ok := cmd.(*Command)

	if !ok {
		return errors.CommandAssertion("Handle")
	}

	id, err := user.NewId(data.Id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	err = handler.Verify.Run(id)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
