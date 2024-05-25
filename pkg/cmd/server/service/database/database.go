package database

import (
	"context"
	"os"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/spersistence"
)

var uri = os.Getenv("DATABASE_URI")

var databaseName = "codexgo"

var Database *spersistence.MongoDB

func Init() error {
	logger.Info("starting mongodb")

	mongoDB, err := spersistence.NewMongoDatabase(uri, databaseName)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	Database = mongoDB

	return nil
}

func Close(ctx context.Context) error {
	logger.Info("closing mongodb")

	err := spersistence.CloseMongoDatabase(ctx, Database)

	if err != nil {
		return serror.BubbleUp(err, "Close")
	}

	return nil
}
