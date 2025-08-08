package mongodb_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/mongodb"
)

type MongoDBTestSuite struct {
	suite.Default
}

func (s *MongoDBTestSuite) TestSentinel() {
	s.Equal(`[A-Za-z0-9]+\.value`, mongodb.RExDuplicateValue)
}

func TestUnitMongoDBSuite(t *testing.T) {
	suite.Run(t, new(MongoDBTestSuite))
}
