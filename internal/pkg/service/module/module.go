package module

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/command"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence"
	"github.com/bastean/codexgo/v4/internal/pkg/service/query"
	"github.com/bastean/codexgo/v4/internal/pkg/service/record/log"
)

var Module = &struct {
	User string
}{
	User: log.Module("User"),
}

func Start() error {
	log.Starting(Module.User)

	collection, err := user.OpenCollection(
		persistence.MongoDB,
		user.CollectionName,
		user.Bcrypt,
	)

	if err != nil {
		log.CannotBeStarted(Module.User)
		return errors.BubbleUp(err, "Start")
	}

	user.Start(
		collection,
		communication.RabbitMQ,
		user.Bcrypt,
	)

	log.Started(Module.User)

	log.Starting(communication.Service.InMemory)

	command.Bus, err = command.NewBus(command.Mapper{
		user.CreateCommandKey: user.CreateHandler,
		user.UpdateCommandKey: user.UpdateHandler,
		user.DeleteCommandKey: user.DeleteHandler,
		user.VerifyCommandKey: user.VerifyHandler,
	})

	if err != nil {
		log.CannotBeStarted(communication.Service.InMemory)
		return errors.BubbleUp(err, "Start")
	}

	query.Bus, err = query.NewBus(query.Mapper{
		user.ReadQueryKey:  user.ReadHandler,
		user.LoginQueryKey: user.LoginHandler,
	})

	if err != nil {
		log.CannotBeStarted(communication.Service.InMemory)
		return errors.BubbleUp(err, "Start")
	}

	log.Started(communication.Service.InMemory)

	return nil
}
