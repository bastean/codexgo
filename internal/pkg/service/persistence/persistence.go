package persistence

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/badgerdb"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
)

var Service = &struct {
	MongoDB, BadgerDB string
}{
	MongoDB:  log.Service("MongoDB"),
	BadgerDB: log.Service("BadgerDB"),
}

var (
	MongoDB  *mongodb.Database
	BadgerDB *badgerdb.Database
	err      error
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
			return errors.BubbleUp(err)
		}

		log.ConnectionEstablishedWith(Service.MongoDB)
	default:
		log.Starting(Service.BadgerDB)

		BadgerDB, err = badgerdb.Open(env.DatabaseBadgerDBDSN)

		if err != nil {
			log.CannotBeStarted(Service.BadgerDB)
			return errors.BubbleUp(err)
		}

		log.Started(Service.BadgerDB)
	}

	return nil
}

func Down(ctx context.Context) error {
	switch {
	case env.HasMongoDB():
		log.ClosingConnectionWith(Service.MongoDB)

		if err = mongodb.Close(ctx, MongoDB); err != nil {
			log.DisconnectionFailedWith(Service.MongoDB)
			return errors.BubbleUp(err)
		}

		log.ConnectionClosedWith(Service.MongoDB)

	default:
		log.Stopping(Service.BadgerDB)

		if err = badgerdb.Close(BadgerDB); err != nil {
			log.CannotBeStopped(Service.BadgerDB)
			return errors.BubbleUp(err)
		}

		log.Stopped(Service.BadgerDB)
	}

	return nil
}
