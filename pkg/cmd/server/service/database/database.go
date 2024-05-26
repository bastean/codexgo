package database

import (
	"context"
	"os"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistences"
)

var uri = os.Getenv("DATABASE_URI")

var databaseName = "codexgo"

var Database *persistences.MongoDB

func Init() error {
	logger.Info("starting mongodb")

	mongoDB, err := persistences.NewMongoDatabase(uri, databaseName)

	if err != nil {
		return errors.BubbleUp(err, "Init")
	}

	Database = mongoDB

	return nil
}

func Close(ctx context.Context) error {
	logger.Info("closing mongodb")

	err := persistences.CloseMongoDatabase(ctx, Database)

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
