package register_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/scommunication"
	"github.com/bastean/codexgo/pkg/context/user/application/register"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserRegisterTestSuite struct {
	suite.Suite
	sut        smodel.CommandHandler[*register.Command]
	useCase    smodel.UseCase[*aggregate.User, *stype.Empty]
	repository *persistence.RepositoryMock
	broker     *scommunication.BrokerMock
}

func (suite *UserRegisterTestSuite) SetupTest() {
	suite.broker = new(scommunication.BrokerMock)
	suite.repository = new(persistence.RepositoryMock)
	suite.useCase = &register.Register{
		Repository: suite.repository,
	}
	suite.sut = &register.CommandHandler{
		UseCase: suite.useCase,
		Broker:  suite.broker,
	}
}

func (suite *UserRegisterTestSuite) TestRegister() {
	command := register.RandomCommand()

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
