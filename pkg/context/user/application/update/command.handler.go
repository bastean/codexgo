package update

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/usecase"
)

type Handler struct {
	usecase.Update
}

func (handler *Handler) Handle(command *Command) error {
	account, err := user.New(&user.Primitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	var updated *user.Password

	if command.UpdatedPassword != "" {
		updated, err = user.NewPassword(command.UpdatedPassword)

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
