package read_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/read"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type ReadTestSuite struct {
	suite.Suite
	sut        handlers.Query[*read.Query, *read.Response]
	read       usecase.Read
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

func (suite *ReadTestSuite) TestRead() {
	random := user.Random()

	query := &read.Query{
		Id: random.Id.Value,
	}

	criteria := &repository.SearchCriteria{
		Id: random.Id,
	}

	suite.repository.On("Search", criteria).Return(random)

	expected := random.ToPrimitive()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitReadSuite(t *testing.T) {
	suite.Run(t, new(ReadTestSuite))
}
