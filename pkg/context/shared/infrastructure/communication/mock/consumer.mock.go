package communicationMock

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/stretchr/testify/mock"
)

type ConsumerMock struct {
	mock.Mock
}

func (m *ConsumerMock) SubscribedTo() []*queue.Queue {
	args := m.Called()
	return args.Get(0).([]*queue.Queue)
}

func (m *ConsumerMock) On(message *message.Message) {
	// TODO?: m.Called(message)
}

func NewConsumerMock() *ConsumerMock {
	return new(ConsumerMock)
}
