package communications

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type CommandBusSuite struct {
	suite.Default
	SUT     roles.CommandBus
	Handler *CommandHandlerMock
	Command *messages.Message
}

func (s *CommandBusSuite) SetupTest() {
	s.Command = messages.Mother().MessageValid()
}

func (s *CommandBusSuite) TestRegister() {
	s.NoError(s.SUT.Register(s.Command.Key, s.Handler))
}

func (s *CommandBusSuite) TestRegisterErrDuplicateCommand() {
	s.NoError(s.SUT.Register(s.Command.Key, s.Handler))

	err := s.SUT.Register(s.Command.Key, s.Handler)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "Register")

	expected := &errors.Internal{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: actual.Where,
		What:  "Already registered",
		Why: errors.Meta{
			"Key": s.Command.Key.Value(),
		},
	}}

	s.Equal(expected, actual)
}

func (s *CommandBusSuite) TestDispatch() {
	s.NoError(s.SUT.Register(s.Command.Key, s.Handler))

	s.Handler.Mock.On("Handle", s.Command)

	s.NoError(s.SUT.Dispatch(s.Command))

	s.Eventually(func() bool {
		return s.Handler.Mock.AssertExpectations(s.T())
	}, 10*time.Second, 30*time.Millisecond)
}

func (s *CommandBusSuite) TestDispatchErrMissingHandler() {
	s.Command = messages.Mother().MessageValid()

	err := s.SUT.Dispatch(s.Command)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	s.Contains(actual.Where, "Dispatch")

	expected := &errors.Internal{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: actual.Where,
		What:  "Failure to execute a Command without a Handler",
		Why: errors.Meta{
			"ID":  s.Command.ID.Value(),
			"Key": s.Command.Key.Value(),
		},
	}}

	s.Equal(expected, actual)
}
