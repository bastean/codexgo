package create

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type CommandHandler struct {
	models.UseCase[*aggregate.User, *types.Empty]
	models.Broker
}

func (handler *CommandHandler) Handle(command *Command) error {
	user, err := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(user)

	if err != nil {
		return errors.BubbleUp(err, "Handle")
	}

	handler.Broker.PublishMessages(user.PullMessages())

	return nil
}
