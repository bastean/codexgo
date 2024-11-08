package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
)

type CommandMock struct {
	mock.Mock
}

func (cmd *CommandMock) Type() commands.Type {
	args := cmd.Called()
	return args.Get(0).(commands.Type)
}

type CommandHandlerMock struct {
	mock.Mock
}

func (handler *CommandHandlerMock) SubscribedTo() commands.Type {
	args := handler.Called()
	return args.Get(0).(commands.Type)
}

func (handler *CommandHandlerMock) Handle(cmd commands.Command) error {
	handler.Called(cmd)
	return nil
}

type CommandBusMock struct {
	mock.Mock
}

func (bus *CommandBusMock) Register(cmd commands.Type, handler commands.Handler) error {
	bus.Called(cmd, handler)
	return nil
}

func (bus *CommandBusMock) Dispatch(cmd commands.Command) error {
	bus.Called(cmd)
	return nil
}
