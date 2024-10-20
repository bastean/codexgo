package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/command"
)

type CommandMock struct {
	mock.Mock
}

func (cmd *CommandMock) Type() command.Type {
	args := cmd.Called()
	return args.Get(0).(command.Type)
}

type CommandHandlerMock struct {
	mock.Mock
}

func (handler *CommandHandlerMock) SubscribedTo() command.Type {
	args := handler.Called()
	return args.Get(0).(command.Type)
}

func (handler *CommandHandlerMock) Handle(cmd command.Command) error {
	handler.Called(cmd)
	return nil
}

type CommandBusMock struct {
	mock.Mock
}

func (bus *CommandBusMock) Register(cmd command.Type, handler command.Handler) error {
	bus.Called(cmd, handler)
	return nil
}

func (bus *CommandBusMock) Dispatch(cmd command.Command) error {
	bus.Called(cmd)
	return nil
}
