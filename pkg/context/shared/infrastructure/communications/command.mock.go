package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
)

type CommandHandlerMock struct {
	mock.Mock
}

func (handler *CommandHandlerMock) Handle(command *commands.Command) error {
	handler.Called(command)
	return nil
}

type CommandBusMock struct {
	mock.Mock
}

func (bus *CommandBusMock) Register(key commands.Key, handler commands.Handler) error {
	bus.Called(key, handler)
	return nil
}

func (bus *CommandBusMock) Dispatch(command *commands.Command) error {
	bus.Called(command)
	return nil
}
