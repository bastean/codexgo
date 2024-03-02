package login_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
	query "github.com/bastean/codexgo/test/pkg/context/user/application/login"
	"github.com/bastean/codexgo/test/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/cryptographic"
	"github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/persistence"
	"github.com/stretchr/testify/suite"
)

type UserLoginTestSuite struct {
	suite.Suite
	repository *persistence.RepositoryMock
	hashing    *cryptographic.HashingMock
	login      *login.Login
	handler    *login.QueryHandler
}

func (suite *UserLoginTestSuite) TestLogin() {
	user := aggregate.Random()

	suite.repository = new(persistence.RepositoryMock)
	suite.hashing = new(cryptographic.HashingMock)
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
