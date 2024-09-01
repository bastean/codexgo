package env

import (
	"os"
)

var (
	DatabaseMongoDBURI  = os.Getenv("CODEXGO_DATABASE_MONGODB_URI")
	DatabaseMongoDBName = os.Getenv("CODEXGO_DATABASE_MONGODB_NAME")
)
