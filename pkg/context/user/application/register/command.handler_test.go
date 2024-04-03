package register_test

import (
	"testing"

	communicationMock "github.com/bastean/codexgo/pkg/context/shared/infrastructure/communication/mock"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	commandMother "github.com/bastean/codexgo/pkg/context/user/application/register/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	persistenceMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/mock"
	"github.com/stretchr/testify/suite"
)

type UserRegisterTestSuite struct {
	suite.Suite
	sut        *register.CommandHandler
	register   *register.Register
	repository *persistenceMock.RepositoryMock
	broker     *communicationMock.BrokerMock
}

func (suite *UserRegisterTestSuite) SetupTest() {
	suite.broker = communicationMock.NewBrokerMock()
	suite.repository = persistenceMock.NewRepositoryMock()
	suite.register = register.NewRegister(suite.repository)
	suite.sut = register.NewCommandHandler(suite.register, suite.broker)
}

func (suite *UserRegisterTestSuite) TestRegister() {
	command := commandMother.Random()

	user := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	messages := user.Messages

	suite.repository.On("Save", user)

	suite.broker.On("PublishMessages", messages)

	suite.sut.Handle(command)

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitUserRegisterSuite(t *testing.T) {
	suite.Run(t, new(UserRegisterTestSuite))
}
