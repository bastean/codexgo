package login_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

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

func (suite *LoginTestSuite) TestSubscribedTo() {
	const expected queries.Type = "user.query.logging.user"

	actual := suite.sut.SubscribedTo()

	suite.Equal(expected, actual)
}

func (suite *LoginTestSuite) TestReplyTo() {
	const expected queries.Type = "user.response.logging.user"

	actual := suite.sut.ReplyTo()

	suite.Equal(expected, actual)
}

func (suite *LoginTestSuite) TestHandle() {
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
