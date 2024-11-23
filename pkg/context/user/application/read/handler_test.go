package read_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
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
	aggregate := user.RandomPrimitive()

	criteria := &repository.SearchCriteria{
		ID: aggregate.ID,
	}

	s.repository.On("Search", criteria).Return(aggregate)

	expected := messages.New[queries.Response](
		read.ResponseKey,
		(*read.ResponseAttributes)(aggregate.ToPrimitive()),
		new(read.ResponseMeta),
	)

	attributes := &read.QueryAttributes{
		ID: aggregate.ID.Value,
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
