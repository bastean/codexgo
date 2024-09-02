package env

import (
	"os"
)

var (
	DatabaseMongoDBURI, DatabaseMongoDBName string
)

func Database() {
	DatabaseMongoDBURI = os.Getenv("CODEXGO_DATABASE_MONGODB_URI")
	DatabaseMongoDBName = os.Getenv("CODEXGO_DATABASE_MONGODB_NAME")
}
