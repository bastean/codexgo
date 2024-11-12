package memory_test

import (
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
	"github.com/stretchr/testify/suite"
)

type CommandBusTestSuite struct {
	suite.Suite
	sut     commands.Bus
	handler *communications.CommandHandlerMock
}

func (s *CommandBusTestSuite) SetupTest() {
	s.handler = new(communications.CommandHandlerMock)

	s.sut = &memory.CommandBus{
		Handlers: make(map[commands.Key]commands.Handler),
	}
}

func (s *CommandBusTestSuite) TestRegister() {
	s.NoError(s.sut.Register(messages.Random[commands.Command]().Key, s.handler))
}

func (s *CommandBusTestSuite) TestRegisterErrDuplicateCommand() {
	key := messages.Random[commands.Command]().Key

	s.NoError(s.sut.Register(key, s.handler))

	err := s.sut.Register(key, s.handler)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Register",
		What:  fmt.Sprintf("%s already registered", key),
		Why: errors.Meta{
			"Command": key,
		},
	}}

	s.EqualError(expected, actual.Error())
}

func (s *CommandBusTestSuite) TestDispatch() {
	command := messages.Random[commands.Command]()

	s.NoError(s.sut.Register(command.Key, s.handler))

	s.handler.On("Handle", command)

	s.NoError(s.sut.Dispatch(command))

	s.handler.AssertExpectations(s.T())
}

func (s *CommandBusTestSuite) TestDispatchErrMissingHandler() {
	command := messages.Random[commands.Command]()

	err := s.sut.Dispatch(command)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Dispatch",
		What:  "Failure to execute a Command without a Handler",
		Why: errors.Meta{
			"Command": command.Key,
		},
	}}

	s.EqualError(expected, actual.Error())
}

func TestIntegrationCommandBusSuite(t *testing.T) {
	suite.Run(t, new(CommandBusTestSuite))
}
