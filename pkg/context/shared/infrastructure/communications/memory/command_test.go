package memory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type CommandBusTestSuite struct {
	suite.Suite
	sut     command.Bus
	command *communications.CommandMock
	handler *communications.CommandHandlerMock
}

func (suite *CommandBusTestSuite) SetupTest() {
	suite.command = new(communications.CommandMock)

	suite.handler = new(communications.CommandHandlerMock)

	suite.sut = &memory.CommandBus{
		Handlers: make(map[command.Type]command.Handler),
	}
}

func (suite *CommandBusTestSuite) TestRegister() {
	const cmd command.Type = "command.testing.register"
	suite.NoError(suite.sut.Register(cmd, suite.handler))
}

func (suite *CommandBusTestSuite) TestRegisterErrDuplicateCommand() {
	const cmd command.Type = "command.testing.register_duplicate"

	suite.NoError(suite.sut.Register(cmd, suite.handler))

	err := suite.sut.Register(cmd, suite.handler)

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Register",
		What:  fmt.Sprintf("%s already registered", cmd),
		Why: errors.Meta{
			"Command": cmd,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CommandBusTestSuite) TestDispatch() {
	const cmd command.Type = "command.testing.dispatch"

	suite.NoError(suite.sut.Register(cmd, suite.handler))

	suite.command.On("Type").Return(cmd)

	suite.handler.On("Handle", suite.command)

	suite.NoError(suite.sut.Dispatch(suite.command))

	suite.command.AssertExpectations(suite.T())

	suite.handler.AssertExpectations(suite.T())
}

func (suite *CommandBusTestSuite) TestDispatchErrMissingHandler() {
	const cmd command.Type = "command.testing.dispatch_missing"

	suite.command.On("Type").Return(cmd)

	err := suite.sut.Dispatch(suite.command)

	suite.command.AssertExpectations(suite.T())

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Dispatch",
		What:  "Failure to execute a Command without a Handler",
		Why: errors.Meta{
			"Command": cmd,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationCommandBusSuite(t *testing.T) {
	suite.Run(t, new(CommandBusTestSuite))
}
