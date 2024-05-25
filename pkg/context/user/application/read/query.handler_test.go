package read_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserReadTestSuite struct {
	suite.Suite
	sut        smodel.QueryHandler[*read.Query, *read.Response]
	useCase    smodel.UseCase[*read.Input, *aggregate.User]
	repository *persistence.RepositoryMock
}

func (suite *UserReadTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.useCase = &read.Read{
		Repository: suite.repository,
	}

	suite.sut = &read.QueryHandler{
		UseCase: suite.useCase,
	}
}

func (suite *UserReadTestSuite) TestLogin() {
	user := aggregate.RandomUser()

	query := &read.Query{
		Id: user.Id.Value(),
	}

	filter := model.RepositorySearchCriteria{
		Id: user.Id,
	}

	suite.repository.On("Search", filter).Return(user)

	expected := user.ToPrimitives()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitUserReadSuite(t *testing.T) {
	suite.Run(t, new(UserReadTestSuite))
}
