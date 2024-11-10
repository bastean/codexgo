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

func (suite *ReadTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.read = &read.Read{
		Repository: suite.repository,
	}

	suite.sut = &read.Handler{
		Read: suite.read,
	}
}

func (suite *ReadTestSuite) TestHandle() {
	account := user.Random()

	criteria := &repository.SearchCriteria{
		Id: account.Id,
	}

	suite.repository.On("Search", criteria).Return(account)

	expected := messages.New[queries.Response](
		read.ResponseKey,
		(*read.ResponseAttributes)(account.ToPrimitive()),
		new(read.ResponseMeta),
	)

	attributes := &read.QueryAttributes{
		Id: account.Id.Value,
	}

	query := messages.RandomWithAttributes[queries.Query](attributes, false)

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitReadSuite(t *testing.T) {
	suite.Run(t, new(ReadTestSuite))
}
