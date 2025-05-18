package env

import (
	"os"
)

var (
	DatabaseMongoDBURI, DatabaseMongoDBName string

	DatabaseBadgerDBDSN string
)

func Database() {
	DatabaseMongoDBURI = os.Getenv(DATABASE_MONGODB_URI)
	DatabaseMongoDBName = os.Getenv(DATABASE_MONGODB_NAME)

	DatabaseBadgerDBDSN = os.Getenv(DATABASE_BADGERDB_DSN)
}

func HasMongoDB() bool {
	return DatabaseMongoDBURI != ""
}
