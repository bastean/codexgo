package register

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type CommandHandler struct {
	model.UseCase[*aggregate.User, *types.Empty]
	model.Broker
}

func (handler *CommandHandler) Handle(command *Command) error {
	user, err := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	if err != nil {
		return errs.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(user)

	if err != nil {
		return errs.BubbleUp(err, "Handle")
	}

	handler.Broker.PublishMessages(user.PullMessages())

	return nil
}
