package badgerdb_test

import (
	"strings"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/badgerdb"
)

type BadgerDBTestSuite struct {
	suite.Default
}

func (s *BadgerDBTestSuite) TestSentinel() {
	s.Equal(":", badgerdb.Separator)
}

func (s *BadgerDBTestSuite) TestNewKey() {
	values := persistences.Mother().KeyValuesValid()

	actual, err := badgerdb.NewKey(values...)

	s.NoError(err)

	expected := []byte(strings.Join(values, ":"))

	s.Equal(expected, actual)
}

func (s *BadgerDBTestSuite) TestParseKey() {
	expected := persistences.Mother().KeyValuesValid()

	key, err := badgerdb.NewKey(expected...)

	s.NoError(err)

	actual := badgerdb.ParseKey(key)

	s.Equal(expected, actual)
}

func TestIntegrationBadgerDBSuite(t *testing.T) {
	suite.Run(t, new(BadgerDBTestSuite))
}
