package communications

import (
	"fmt"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type CommandBusSuite struct {
	suite.Suite
	SUT     roles.CommandBus
	Handler *CommandHandlerMock
}

func (s *CommandBusSuite) TestRegister() {
	s.NoError(s.SUT.Register(messages.Random().Key, s.Handler))
}

func (s *CommandBusSuite) TestRegisterErrDuplicateCommand() {
	key := messages.Random().Key

	s.NoError(s.SUT.Register(key, s.Handler))

	err := s.SUT.Register(key, s.Handler)

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

	s.Equal(expected, actual)
}

func (s *CommandBusSuite) TestDispatch() {
	command := messages.Random()

	s.NoError(s.SUT.Register(command.Key, s.Handler))

	s.Handler.Mock.On("Handle", command)

	s.NoError(s.SUT.Dispatch(command))

	s.Handler.Mock.AssertExpectations(s.T())
}

func (s *CommandBusSuite) TestDispatchErrMissingHandler() {
	command := messages.Random()

	err := s.SUT.Dispatch(command)

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

	s.Equal(expected, actual)
}
