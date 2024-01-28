package container

import (
	"github.com/bastean/codexgo/context/pkg/shared/infrastructure/persistence"
)

var Database = persistence.NewMongoDatabase()
