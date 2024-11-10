package login_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type LoginTestSuite struct {
	suite.Suite
	sut        queries.Handler
	login      cases.Login
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

func (suite *LoginTestSuite) TestHandle() {
	account := user.Random()

	criteria := &repository.SearchCriteria{
		Email: account.Email,
	}

	suite.repository.On("Search", criteria).Return(account)

	suite.hashing.On("IsNotEqual", account.Password.Value, account.Password.Value).Return(false)

	expected := messages.New[queries.Response](
		login.ResponseKey,
		(*login.ResponseAttributes)(account.ToPrimitive()),
		new(login.ResponseMeta),
	)

	attributes := &login.QueryAttributes{
		Email:    account.Email.Value,
		Password: account.Password.Value,
	}

	query := messages.RandomWithAttributes[queries.Query](attributes, false)

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
