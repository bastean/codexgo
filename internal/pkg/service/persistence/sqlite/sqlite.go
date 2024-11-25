package sqlite

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/sqlite"
)

type (
	Database = sqlite.Database
)

var (
	Open  = sqlite.Open
	Close = sqlite.Close
)
