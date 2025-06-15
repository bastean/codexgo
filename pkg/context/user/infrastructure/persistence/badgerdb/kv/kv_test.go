package kv_test

import (
	"os"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/badgerdb"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/badgerdb/kv"
)

type KVTestSuite struct {
	persistence.RepositorySuite
}

func (s *KVTestSuite) SetupSuite() {
	s.NoError(os.RemoveAll(os.Getenv("CODEXGO_DATABASE_BADGERDB_DSN")))

	session, err := badgerdb.Open(os.Getenv("CODEXGO_DATABASE_BADGERDB_DSN"))

	if err != nil {
		errors.Panic(err)
	}

	s.RepositorySuite.SUT, err = kv.Open(session)

	if err != nil {
		errors.Panic(err)
	}
}

func (s *KVTestSuite) TestSentinel() {
	s.Equal(3, kv.TotalCriteriaFields)
	s.Equal(0, kv.IxID)
	s.Equal(1, kv.IxEmail)
	s.Equal(2, kv.IxUsername)
}

func (s *KVTestSuite) TearDownSuite() {
	s.NoError(os.RemoveAll(os.Getenv("CODEXGO_DATABASE_BADGERDB_DSN")))
}

func TestIntegrationKVSuite(t *testing.T) {
	suite.Run(t, new(KVTestSuite))
}
