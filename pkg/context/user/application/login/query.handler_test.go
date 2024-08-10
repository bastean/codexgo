package login_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	sut        handlers.Query[*login.Query, *login.Response]
	login      usecase.Login
	hashing    *cryptographic.HashingMock
	repository *persistence.UserMock
}

func (suite *LoginTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.login = &login.Login{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &login.Handler{
		Login: suite.login,
	}
}

func (suite *LoginTestSuite) TestLogin() {
	random := user.Random()

	query := &login.Query{
		Email:    random.Email.Value,
		Password: random.Password.Value,
	}

	criteria := &repository.SearchCriteria{
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

func TestUnitLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
