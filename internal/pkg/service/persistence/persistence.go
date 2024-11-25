package persistence

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence/mongodb"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence/sqlite"
	"github.com/bastean/codexgo/v4/internal/pkg/service/record/log"
)

var Service = &struct {
	MongoDB, SQLite string
}{
	MongoDB: log.Service("MongoDB"),
	SQLite:  log.Service("SQLite"),
}

var (
	err     error
	MongoDB *mongodb.Database
	SQLite  *sqlite.Database
)

func Up() error {
	switch {
	case env.HasDatabase():
		log.EstablishingConnectionWith(Service.MongoDB)

		MongoDB, err = mongodb.Open(
			env.DatabaseMongoDBURI,
			env.DatabaseMongoDBName,
		)

		if err != nil {
			log.ConnectionFailedWith(Service.MongoDB)
			return errors.BubbleUp(err, "Up")
		}

		log.ConnectionEstablishedWith(Service.MongoDB)
	default:
		log.Starting(Service.SQLite)

		SQLite, err = sqlite.Open(env.DatabaseSQLiteDSN)

		if err != nil {
			log.CannotBeStarted(Service.SQLite)
			return errors.BubbleUp(err, "Up")
		}

		log.Started(Service.SQLite)
	}

	return nil
}

func Down(ctx context.Context) error {
	switch {
	case env.HasDatabase():
		log.ClosingConnectionWith(Service.MongoDB)

		if err = mongodb.Close(ctx, MongoDB); err != nil {
			log.DisconnectionFailedWith(Service.MongoDB)
			return errors.BubbleUp(err, "Down")
		}

		log.ConnectionClosedWith(Service.MongoDB)

	default:
		log.Stopping(Service.SQLite)

		if err = sqlite.Close(SQLite); err != nil {
			log.CannotBeStopped(Service.SQLite)
			return errors.BubbleUp(err, "Down")
		}

		log.Stopped(Service.SQLite)
	}

	return nil
}
