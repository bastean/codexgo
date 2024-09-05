package env

import (
	"os"
)

var (
	DatabaseMongoDBURI, DatabaseMongoDBName string
)

func Database() {
	DatabaseMongoDBURI = os.Getenv(DATABASE_MONGODB_URI)
	DatabaseMongoDBName = os.Getenv(DATABASE_MONGODB_NAME)
}
