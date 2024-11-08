package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

type EventConsumerMock struct {
	mock.Mock
}

func (m *EventConsumerMock) On(event *events.Event) error {
	m.Called(event)
	return nil
}

type EventBusMock struct {
	mock.Mock
}

func (m *EventBusMock) Subscribe(key events.Key, consumers events.Consumer) error {
	m.Called(key, consumers)
	return nil
}

func (m *EventBusMock) Publish(event *events.Event) error {
	m.Called(event)
	return nil
}
