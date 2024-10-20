package create_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type CreateTestSuite struct {
	suite.Suite
	sut        command.Handler
	create     cases.Create
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

func (suite *CreateTestSuite) TestSubscribedTo() {
	const expected command.Type = "user.command.creating.user"

	actual := suite.sut.SubscribedTo()

	suite.Equal(expected, actual)
}

func (suite *CreateTestSuite) TestHandle() {
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
