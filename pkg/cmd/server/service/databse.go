package service

import (
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistence/database"
)

var Database = database.NewMongoDatabase()
