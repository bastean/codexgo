package database

import (
	"context"
	"os"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistence/database"
)

var uri = os.Getenv("DATABASE_URI")

var databaseName = "codexgo"

var Database = database.NewMongoDatabase(uri, databaseName)

func Init() {
	logger.Logger.Info("starting mongodb")
}

func Close(ctx context.Context) {
	database.CloseMongoDatabase(ctx, Database)
	logger.Logger.Info("mongodb closed")
}
