package read_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/application/read"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type ReadHandlerTestSuite struct {
	suite.Suite
	sut        models.QueryHandler[*read.Query, *read.Response]
	usecase    models.UseCase[models.ValueObject[string], *aggregate.User]
	repository *persistence.RepositoryMock
}

func (suite *ReadHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.usecase = &read.Read{
		Repository: suite.repository,
	}

	suite.sut = &read.Handler{
		UseCase: suite.usecase,
	}
}

func (suite *ReadHandlerTestSuite) TestLogin() {
	user := aggregate.RandomUser()

	query := &read.Query{
		Id: user.Id.Value(),
	}

	criteria := &model.RepositorySearchCriteria{
		Id: user.Id,
	}

	suite.repository.On("Search", criteria).Return(user)

	expected := user.ToPrimitives()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitReadHandlerSuite(t *testing.T) {
	suite.Run(t, new(ReadHandlerTestSuite))
}
