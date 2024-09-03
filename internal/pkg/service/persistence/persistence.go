package persistence

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/persistence/mongodb"
)

var Service = &struct {
	MongoDB string
}{
	MongoDB: log.Service("MongoDB"),
}

var (
	err     error
	MongoDB *mongodb.MongoDB
)

func Up() error {
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

	return nil
}

func Down(ctx context.Context) error {
	log.ClosingConnectionWith(Service.MongoDB)

	if err = mongodb.Close(ctx, MongoDB); err != nil {
		log.DisconnectionFailedWith(Service.MongoDB)
		return errors.BubbleUp(err, "Down")
	}

	log.ConnectionClosedWith(Service.MongoDB)

	return nil
}
