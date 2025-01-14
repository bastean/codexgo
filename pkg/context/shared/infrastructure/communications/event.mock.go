package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type EventConsumerMock struct {
	mock.Mock
}

func (m *EventConsumerMock) On(event *messages.Message) error {
	m.Called(event)
	return nil
}

type EventBusMock struct {
	mock.Mock
}

func (m *EventBusMock) Subscribe(key messages.Key, consumers roles.EventConsumer) error {
	m.Called(key, consumers)
	return nil
}

func (m *EventBusMock) Publish(event *messages.Message) error {
	m.Called(event)
	return nil
}
