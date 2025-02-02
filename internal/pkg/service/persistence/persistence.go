package persistence

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/sqlite"
)

var Service = &struct {
	MongoDB, SQLite string
}{
	MongoDB: log.Service("MongoDB"),
	SQLite:  log.Service("SQLite"),
}

var (
	MongoDB *mongodb.Database
	SQLite  *sqlite.Database
	err     error
)

func Up() error {
	switch {
	case env.HasMongoDB():
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
	case env.HasMongoDB():
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
