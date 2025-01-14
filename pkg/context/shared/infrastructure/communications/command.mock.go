package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type CommandHandlerMock struct {
	mock.Mock
}

func (m *CommandHandlerMock) Handle(command *messages.Message) error {
	m.Called(command)
	return nil
}

type CommandBusMock struct {
	mock.Mock
}

func (m *CommandBusMock) Register(key messages.Key, handler roles.CommandHandler) error {
	m.Called(key, handler)
	return nil
}

func (m *CommandBusMock) Dispatch(command *messages.Message) error {
	m.Called(command)
	return nil
}
