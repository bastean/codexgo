package update

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
)

type Input struct {
	User            *aggregate.User
	UpdatedPassword models.ValueObject[string]
}

type Handler struct {
	models.UseCase[*Input, types.Empty]
}

func (handler *Handler) Handle(command *Command) error {
	user, err := aggregate.NewUser(&aggregate.UserPrimitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	var updatedPassword models.ValueObject[string]

	if command.UpdatedPassword != "" {
		updatedPassword, err = valueobj.NewPassword(command.UpdatedPassword)

		if err != nil {
			return errors.BubbleUp(err, "Handle")
		}
	}

	_, err = handler.UseCase.Run(&Input{
		User:            user,
		UpdatedPassword: updatedPassword,
	})

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	return nil
}
