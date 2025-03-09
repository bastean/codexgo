package table_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/persistences/sqlite"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence/sqlite/table"
)

type TableTestSuite struct {
	persistence.RepositorySuite
}

func (s *TableTestSuite) SetupSuite() {
	_ = os.Remove(os.Getenv("CODEXGO_DATABASE_SQLITE_DSN"))

	session, err := sqlite.Open(os.Getenv("CODEXGO_DATABASE_SQLITE_DSN"))

	if err != nil {
		errors.Panic(err.Error(), "SetupSuite")
	}

	s.RepositorySuite.SUT, err = table.Open(session)

	if err != nil {
		errors.Panic(err.Error(), "SetupSuite")
	}
}

func (*TableTestSuite) TearDownSuite() {
	_ = os.Remove(os.Getenv("CODEXGO_DATABASE_SQLITE_DSN"))
}

func TestIntegrationTableSuite(t *testing.T) {
	suite.Run(t, new(TableTestSuite))
}
