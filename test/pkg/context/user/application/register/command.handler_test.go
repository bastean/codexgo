package register_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/application/register"
	commandMother "github.com/bastean/codexgo/test/pkg/context/user/application/register"
	aggregateMother "github.com/bastean/codexgo/test/pkg/context/user/domain/aggregate"
	persistenceMock "github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/persistence"
	"github.com/stretchr/testify/suite"
)

type UserRegisterTestSuite struct {
	suite.Suite
	sut        *register.CommandHandler
	register   *register.Register
	repository *persistenceMock.RepositoryMock
}

func (suite *UserRegisterTestSuite) SetupTest() {
	suite.repository = persistenceMock.NewRepositoryMock()
	suite.register = &register.Register{Repository: suite.repository}
	suite.register = register.NewRegister(suite.repository)
	suite.sut = register.NewCommandHandler(suite.register)
}

func (suite *UserRegisterTestSuite) TestRegister() {
	command := commandMother.Random()

	user := aggregateMother.FromCommand(*command)

	suite.repository.On("Save", user)

	suite.sut.Handle(command)

	suite.repository.AssertExpectations(suite.T())
}

func TestUserRegisterSuite(t *testing.T) {
	suite.Run(t, new(UserRegisterTestSuite))
}
