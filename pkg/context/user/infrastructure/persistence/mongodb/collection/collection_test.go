package collection_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/mongodb/collection"
)

type CollectionTestSuite struct {
	persistence.RepositorySuite
}

func (s *CollectionTestSuite) SetupTest() {
	session, err := mongodb.Open(
		os.Getenv("CODEXGO_DATABASE_MONGODB_URI"),
		os.Getenv("CODEXGO_DATABASE_MONGODB_NAME"),
	)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}

	name := "users-test"

	s.RepositorySuite.SUT, err = collection.Open(session, name)

	if err != nil {
		errors.Panic(err.Error(), "SetupTest")
	}
}

func TestIntegrationCollectionSuite(t *testing.T) {
	suite.Run(t, new(CollectionTestSuite))
}
