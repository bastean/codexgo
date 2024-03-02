package service

import (
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/persistence"
)

var Database = persistence.NewMongoDatabase()
