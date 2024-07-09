package login_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type LoginHandlerTestSuite struct {
	suite.Suite
	sut        handlers.Query[*login.Query, *login.Response]
	login      usecase.Login
	hashing    *cryptographic.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *LoginHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.login = &login.Login{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &login.Handler{
		Login: suite.login,
	}
}

func (suite *LoginHandlerTestSuite) TestLogin() {
	random := user.Random()

	query := &login.Query{
		Email:    random.Email.Value,
		Password: random.Password.Value,
	}

	criteria := &model.RepositorySearchCriteria{
		Email: random.Email,
	}

	suite.repository.On("Search", criteria).Return(random)

	suite.hashing.On("IsNotEqual", random.Password.Value, random.Password.Value).Return(false)

	expected := random.ToPrimitive()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitLoginHandlerSuite(t *testing.T) {
	suite.Run(t, new(LoginHandlerTestSuite))
}
