package login_test

import (
	"testing"

	"github.com/bastean/codexgo/context/pkg/user/application/login"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"github.com/bastean/codexgo/test/contexts/crud/user/__mocks__/infrastructure/cryptographic"
	"github.com/bastean/codexgo/test/contexts/crud/user/__mocks__/infrastructure/persistence"
	query "github.com/bastean/codexgo/test/contexts/crud/user/application/login"
	"github.com/bastean/codexgo/test/contexts/crud/user/domain/aggregate"
	"github.com/stretchr/testify/suite"
)

type UserLoginTestSuite struct {
	suite.Suite
	repository *persistence.UserRepositoryMock
	hashing    *cryptographic.UserHashingMock
	login      *login.Login
	handler    *login.QueryHandler
}

func (suite *UserLoginTestSuite) TestLogin() {
	user := aggregate.Random()

	suite.repository = new(persistence.UserRepositoryMock)
	suite.hashing = new(cryptographic.UserHashingMock)
	suite.login = &login.Login{Repository: suite.repository, Hashing: suite.hashing}
	suite.handler = &login.QueryHandler{Login: *suite.login}

	query := query.Create(user.Email, user.Password)

	filter := repository.Filter{Email: user.Email}

	suite.repository.On("Search", filter).Return(user)

	response := suite.handler.Handle(*query)

	suite.repository.AssertCalled(suite.T(), "Search", filter)

	suite.EqualValues(user.ToPrimitives(), response)
}

func TestUserLoginSuite(t *testing.T) {
	suite.Run(t, new(UserLoginTestSuite))
}
