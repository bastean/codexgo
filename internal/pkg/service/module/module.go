package module

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/cipher/bcrypt"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/command"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/query"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence"
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
		bcrypt.Bcrypt,
	)

	if err != nil {
		log.CannotBeStarted(Module.User)
		return errors.BubbleUp(err, "Start")
	}

	user.Start(
		collection,
		communication.Bus,
		bcrypt.Bcrypt,
	)

	log.Started(Module.User)

	log.Starting(communication.Service.CommandBus)

	command.Bus, err = command.NewBus(command.Mapper{
		user.CreateCommandKey: user.CreateHandler,
		user.UpdateCommandKey: user.UpdateHandler,
		user.DeleteCommandKey: user.DeleteHandler,
		user.VerifyCommandKey: user.VerifyHandler,
	})

	if err != nil {
		log.CannotBeStarted(communication.Service.CommandBus)
		return errors.BubbleUp(err, "Start")
	}

	log.Started(communication.Service.CommandBus)

	log.Starting(communication.Service.QueryBus)

	query.Bus, err = query.NewBus(query.Mapper{
		user.ReadQueryKey:  user.ReadHandler,
		user.LoginQueryKey: user.LoginHandler,
	})

	if err != nil {
		log.CannotBeStarted(communication.Service.QueryBus)
		return errors.BubbleUp(err, "Start")
	}

	log.Started(communication.Service.QueryBus)

	return nil
}
