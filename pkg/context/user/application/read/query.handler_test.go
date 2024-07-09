package read_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type ReadHandlerTestSuite struct {
	suite.Suite
	sut        handlers.Query[*read.Query, *read.Response]
	read       usecase.Read
	repository *persistence.RepositoryMock
}

func (suite *ReadHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.read = &read.Read{
		Repository: suite.repository,
	}

	suite.sut = &read.Handler{
		Read: suite.read,
	}
}

func (suite *ReadHandlerTestSuite) TestLogin() {
	random := user.Random()

	query := &read.Query{
		Id: random.Id.Value,
	}

	criteria := &model.RepositorySearchCriteria{
		Id: random.Id,
	}

	suite.repository.On("Search", criteria).Return(random)

	expected := random.ToPrimitive()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitReadHandlerSuite(t *testing.T) {
	suite.Run(t, new(ReadHandlerTestSuite))
}
