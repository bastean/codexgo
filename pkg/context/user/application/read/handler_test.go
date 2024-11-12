package read_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type ReadTestSuite struct {
	suite.Suite
	sut        queries.Handler
	read       cases.Read
	repository *persistence.UserMock
}

func (s *ReadTestSuite) SetupTest() {
	s.repository = new(persistence.UserMock)

	s.read = &read.Case{
		Repository: s.repository,
	}

	s.sut = &read.Handler{
		Read: s.read,
	}
}

func (s *ReadTestSuite) TestHandle() {
	account := user.Random()

	criteria := &repository.SearchCriteria{
		ID: account.ID,
	}

	s.repository.On("Search", criteria).Return(account)

	expected := messages.New[queries.Response](
		read.ResponseKey,
		(*read.ResponseAttributes)(account.ToPrimitive()),
		new(read.ResponseMeta),
	)

	attributes := &read.QueryAttributes{
		ID: account.ID.Value,
	}

	query := messages.RandomWithAttributes[queries.Query](attributes, false)

	actual, err := s.sut.Handle(query)

	s.NoError(err)

	s.repository.AssertExpectations(s.T())

	s.EqualValues(expected, actual)
}

func TestUnitReadSuite(t *testing.T) {
	suite.Run(t, new(ReadTestSuite))
}
