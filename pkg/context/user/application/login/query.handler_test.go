package login_test

import (
	"testing"

	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	aggregateMother "github.com/bastean/codexgo/pkg/context/user/domain/aggregate/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	cryptographicMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic/mock"
	persistenceMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/mock"
	"github.com/stretchr/testify/suite"
)

type UserLoginTestSuite struct {
	suite.Suite
	sut        sharedModel.QueryHandler[*login.Query, *login.Response]
	login      sharedModel.UseCase[*login.Input, *aggregate.User]
	hashing    *cryptographicMock.HashingMock
	repository *persistenceMock.RepositoryMock
}

func (suite *UserLoginTestSuite) SetupTest() {
	suite.repository = new(persistenceMock.RepositoryMock)
	suite.hashing = new(cryptographicMock.HashingMock)
	suite.login = &login.Login{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}
	suite.sut = &login.QueryHandler{
		UseCase: suite.login,
	}
}

func (suite *UserLoginTestSuite) TestLogin() {
	user := aggregateMother.Random()

	query := &login.Query{
		Email:    user.Email.Value(),
		Password: user.Password.Value(),
	}

	filter := model.RepositorySearchCriteria{
		Email: user.Email,
	}

	suite.repository.On("Search", filter).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value, user.Password.Value).Return(false)

	expected := user.ToPrimitives()

	actual, err := suite.sut.Handle(query)

	suite.NoError(err)

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())

	suite.EqualValues(expected, actual)
}

func TestUnitUserLoginSuite(t *testing.T) {
	suite.Run(t, new(UserLoginTestSuite))
}
