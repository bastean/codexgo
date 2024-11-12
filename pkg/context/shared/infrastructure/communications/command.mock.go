package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
)

type CommandHandlerMock struct {
	mock.Mock
}

func (m *CommandHandlerMock) Handle(command *commands.Command) error {
	m.Called(command)
	return nil
}

type CommandBusMock struct {
	mock.Mock
}

func (m *CommandBusMock) Register(key commands.Key, handler commands.Handler) error {
	m.Called(key, handler)
	return nil
}

func (m *CommandBusMock) Dispatch(command *commands.Command) error {
	m.Called(command)
	return nil
}
