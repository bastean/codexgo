package create_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type CreateTestSuite struct {
	suite.Suite
	sut        commands.Handler
	create     cases.Create
	repository *persistence.UserMock
	bus        *communications.EventBusMock
}

func (suite *CreateTestSuite) SetupTest() {
	suite.bus = new(communications.EventBusMock)

	suite.repository = new(persistence.UserMock)

	suite.create = &create.Create{
		Repository: suite.repository,
	}

	suite.sut = &create.Handler{
		Create: suite.create,
		Bus:    suite.bus,
	}
}

func (suite *CreateTestSuite) TestSubscribedTo() {
	const expected commands.Type = "user.command.creating.user"

	actual := suite.sut.SubscribedTo()

	suite.Equal(expected, actual)
}

func (suite *CreateTestSuite) TestHandle() {
	command := create.RandomCommand()

	account, err := user.New(&user.Primitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.Password,
	})

	suite.NoError(err)

	suite.repository.On("Create", account)

	for _, event := range account.Events {
		suite.bus.On("Publish", event)
	}

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.bus.AssertExpectations(suite.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
