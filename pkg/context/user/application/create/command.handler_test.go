package create_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/pkg/context/user/application/create"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateHandlerTestSuite struct {
	suite.Suite
	sut        handlers.Command[*create.Command]
	create     usecase.Create
	repository *persistence.UserMock
	broker     *communications.BrokerMock
}

func (suite *CreateHandlerTestSuite) SetupTest() {
	suite.broker = new(communications.BrokerMock)

	suite.repository = new(persistence.UserMock)

	suite.create = &create.Create{
		User: suite.repository,
	}

	suite.sut = &create.Handler{
		Create: suite.create,
		Broker: suite.broker,
	}
}

func (suite *CreateHandlerTestSuite) TestCreate() {
	command := create.RandomCommand()

	new, _ := user.New(&user.Primitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
	})

	messages := new.Messages

	suite.repository.On("Save", new)

	suite.broker.On("PublishMessages", messages)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.broker.AssertExpectations(suite.T())
}

func TestUnitCreateHandlerSuite(t *testing.T) {
	suite.Run(t, new(CreateHandlerTestSuite))
}
