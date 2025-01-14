package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type QueryHandlerMock struct {
	mock.Mock
}

func (m *QueryHandlerMock) Handle(query *messages.Message) (*messages.Message, error) {
	args := m.Called(query)
	return args.Get(0).(*messages.Message), nil
}

type QueryBusMock struct {
	mock.Mock
}

func (m *QueryBusMock) Register(key messages.Key, handler roles.QueryHandler) error {
	m.Called(key, handler)
	return nil
}

func (m *QueryBusMock) Ask(query *messages.Message) (*messages.Message, error) {
	args := m.Called(query)
	return args.Get(0).(*messages.Message), nil
}
