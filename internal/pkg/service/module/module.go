package module

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence"
)

var Module = &struct {
	User string
}{
	User: log.Module("User"),
}

func Up() error {
	log.Starting(Module.User)

	collection, err := user.OpenCollection(
		persistence.MongoDB,
		user.CollectionName,
		user.Bcrypt,
	)

	if err != nil {
		log.CannotBeStarted(Module.User)
		return errors.BubbleUp(err, "Up")
	}

	user.Start(
		collection,
		communication.RabbitMQ,
		user.Bcrypt,
	)

	log.Started(Module.User)

	return nil
}
