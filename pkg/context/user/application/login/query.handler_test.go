package login_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type LoginHandlerTestSuite struct {
	suite.Suite
	sut        models.QueryHandler[*login.Query, *login.Response]
	usecase    models.UseCase[*login.Input, *aggregate.User]
	hashing    *cryptographic.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *LoginHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.usecase = &login.Login{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &login.Handler{
		UseCase: suite.usecase,
	}
}

func (suite *LoginHandlerTestSuite) TestLogin() {
	user := aggregate.RandomUser()

	query := &login.Query{
		Email:    user.Email.Value(),
		Password: user.Password.Value(),
	}

	criteria := &model.RepositorySearchCriteria{
		Email: user.Email,
	}

	suite.repository.On("Search", criteria).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value(), user.Password.Value()).Return(false)

	expected := user.ToPrimitives()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitLoginHandlerSuite(t *testing.T) {
	suite.Run(t, new(LoginHandlerTestSuite))
}
