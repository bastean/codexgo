package env

import (
	"os"
)

var (
	DatabaseMongoDBURI  = os.Getenv("DATABASE_MONGODB_URI")
	DatabaseMongoDBName = os.Getenv("DATABASE_MONGODB_NAME")
)
