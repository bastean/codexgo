package database

import (
	"context"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistence/database"
)

var Database = database.NewMongoDatabase()

func Init() {
	logger.Logger.Info("starting mongodb")
}

func Close(ctx context.Context) {
	database.CloseMongoDatabase(ctx, Database)
	logger.Logger.Info("mongodb closed")
}
