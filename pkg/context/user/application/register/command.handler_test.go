package register_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	communicationMock "github.com/bastean/codexgo/pkg/context/shared/infrastructure/communication/mock"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	commandMother "github.com/bastean/codexgo/pkg/context/user/application/register/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	persistenceMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/mock"
	"github.com/stretchr/testify/suite"
)

type UserRegisterTestSuite struct {
	suite.Suite
	sut        model.CommandHandler[*register.Command]
	register   model.UseCase[*aggregate.User, *types.Empty]
	repository *persistenceMock.RepositoryMock
	broker     *communicationMock.BrokerMock
}

func (suite *UserRegisterTestSuite) SetupTest() {
	suite.broker = new(communicationMock.BrokerMock)
	suite.repository = new(persistenceMock.RepositoryMock)
	suite.register = &register.Register{
		Repository: suite.repository,
	}
	suite.sut = &register.CommandHandler{
		UseCase: suite.register,
		Broker:  suite.broker,
	}
}

func (suite *UserRegisterTestSuite) TestRegister() {
	command := commandMother.Random()

	user, _ := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	messages := user.Messages

	suite.repository.On("Save", user)

	suite.broker.On("PublishMessages", messages)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitUserRegisterSuite(t *testing.T) {
	suite.Run(t, new(UserRegisterTestSuite))
}
