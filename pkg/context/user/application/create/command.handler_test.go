package create_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateTestSuite struct {
	suite.Suite
	sut        handlers.Command[*create.Command]
	create     usecase.Create
	repository *persistence.UserMock
	broker     *communications.BrokerMock
}

func (suite *CreateTestSuite) SetupTest() {
	suite.broker = new(communications.BrokerMock)

	suite.repository = new(persistence.UserMock)

	suite.create = &create.Create{
		Repository: suite.repository,
	}

	suite.sut = &create.Handler{
		Create: suite.create,
		Broker: suite.broker,
	}
}

func (suite *CreateTestSuite) TestCreate() {
	command := create.RandomCommand()

	new, err := user.New(&user.Primitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
	})

	suite.NoError(err)

	messages := new.Messages

	suite.repository.On("Create", new)

	suite.broker.On("PublishMessages", messages)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
