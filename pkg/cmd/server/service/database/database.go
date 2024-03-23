package database

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistence/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var Database *mongo.Database

func init() {
	logger.Logger.Info("starting mongodb")

	Database = database.NewMongoDatabase()
}
