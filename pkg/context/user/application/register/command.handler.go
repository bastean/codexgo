package register

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
)

type CommandHandler struct {
	smodel.UseCase[*aggregate.User, *stype.Empty]
	smodel.Broker
}

func (handler *CommandHandler) Handle(command *Command) error {
	user, err := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	if err != nil {
		return serror.BubbleUp(err, "Handle")
	}

	_, err = handler.UseCase.Run(user)

	if err != nil {
		return serror.BubbleUp(err, "Handle")
	}

	handler.Broker.PublishMessages(user.PullMessages())

	return nil
}
