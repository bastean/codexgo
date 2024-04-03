package login_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/application/login"
	aggregateMother "github.com/bastean/codexgo/pkg/context/user/domain/aggregate/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	cryptographicMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic/mock"
	persistenceMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/mock"
	"github.com/stretchr/testify/suite"
)

type UserLoginTestSuite struct {
	suite.Suite
	sut        *login.QueryHandler
	login      *login.Login
	hashing    *cryptographicMock.HashingMock
	repository *persistenceMock.RepositoryMock
}

func (suite *UserLoginTestSuite) SetupTest() {
	suite.repository = persistenceMock.NewRepositoryMock()
	suite.hashing = cryptographicMock.NewHashingMock()
	suite.login = login.NewLogin(suite.repository, suite.hashing)
	suite.sut = login.NewQueryHandler(suite.login)
}

func (suite *UserLoginTestSuite) TestLogin() {
	user := aggregateMother.Random()

	query := login.NewQuery(user.Email.Value, user.Password.Value)

	filter := model.RepositorySearchFilter{Email: user.Email}

	suite.repository.On("Search", filter).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value, user.Password.Value).Return(false)

	expected := user.ToPrimitives()

	actual := suite.sut.Handle(query)

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitUserLoginSuite(t *testing.T) {
	suite.Run(t, new(UserLoginTestSuite))
}
