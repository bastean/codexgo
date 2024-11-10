package memory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type CommandBusTestSuite struct {
	suite.Suite
	sut     commands.Bus
	handler *communications.CommandHandlerMock
}

func (suite *CommandBusTestSuite) SetupTest() {
	suite.handler = new(communications.CommandHandlerMock)

	suite.sut = &memory.CommandBus{
		Handlers: make(map[commands.Key]commands.Handler),
	}
}

func (suite *CommandBusTestSuite) TestRegister() {
	suite.NoError(suite.sut.Register(messages.Random[commands.Command]().Key, suite.handler))
}

func (suite *CommandBusTestSuite) TestRegisterErrDuplicateCommand() {
	key := messages.Random[commands.Command]().Key

	suite.NoError(suite.sut.Register(key, suite.handler))

	err := suite.sut.Register(key, suite.handler)

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Register",
		What:  fmt.Sprintf("%s already registered", key),
		Why: errors.Meta{
			"Command": key,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CommandBusTestSuite) TestDispatch() {
	command := messages.Random[commands.Command]()

	suite.NoError(suite.sut.Register(command.Key, suite.handler))

	suite.handler.On("Handle", command)

	suite.NoError(suite.sut.Dispatch(command))

	suite.handler.AssertExpectations(suite.T())
}

func (suite *CommandBusTestSuite) TestDispatchErrMissingHandler() {
	command := messages.Random[commands.Command]()

	err := suite.sut.Dispatch(command)

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Dispatch",
		What:  "Failure to execute a Command without a Handler",
		Why: errors.Meta{
			"Command": command.Key,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationCommandBusSuite(t *testing.T) {
	suite.Run(t, new(CommandBusTestSuite))
}
