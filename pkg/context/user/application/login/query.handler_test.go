package login_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/user/application/login"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserLoginTestSuite struct {
	suite.Suite
	sut        smodel.QueryHandler[*login.Query, *login.Response]
	useCase    smodel.UseCase[*login.Input, *aggregate.User]
	hashing    *cryptographic.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *UserLoginTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.hashing = new(cryptographic.HashingMock)
	suite.useCase = &login.Login{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}
	suite.sut = &login.QueryHandler{
		UseCase: suite.useCase,
	}
}

func (suite *UserLoginTestSuite) TestLogin() {
	user := aggregate.RandomUser()

	query := &login.Query{
		Email:    user.Email.Value(),
		Password: user.Password.Value(),
	}

	filter := model.RepositorySearchCriteria{
		Email: user.Email,
	}

	suite.repository.On("Search", filter).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value(), user.Password.Value()).Return(false)

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
