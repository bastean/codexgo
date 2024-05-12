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
	database, err := spersistence.NewMongoDatabase(uri, databaseName)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	Database = database

	logger.Info("starting mongodb")

	return nil
}

func Close(ctx context.Context) error {
	err := spersistence.CloseMongoDatabase(ctx, Database)

	if err != nil {
		return serror.BubbleUp(err, "Close")
	}

	logger.Info("mongodb closed")

	return nil
}
