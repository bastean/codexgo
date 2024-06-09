package create_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserCreateTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*create.Command]
	useCase    models.UseCase[*aggregate.User, types.Empty]
	repository *persistence.RepositoryMock
	broker     *communications.BrokerMock
}

func (suite *UserCreateTestSuite) SetupTest() {
	suite.broker = new(communications.BrokerMock)

	suite.repository = new(persistence.RepositoryMock)

	suite.useCase = &create.Create{
		Repository: suite.repository,
	}

	suite.sut = &create.CommandHandler{
		UseCase: suite.useCase,
		Broker:  suite.broker,
	}
}

func (suite *UserCreateTestSuite) TestCreate() {
	command := create.RandomCommand()

	user, _ := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	messages := user.Messages

	suite.repository.On("Save", user)

	suite.broker.On("PublishMessages", messages)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitUserCreateSuite(t *testing.T) {
	suite.Run(t, new(UserCreateTestSuite))
}
