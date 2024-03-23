package communicationMock

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/exchange"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/stretchr/testify/mock"
)

type BrokerMock struct {
	mock.Mock
}

func (m *BrokerMock) PublishMessages(messages []*message.Message) {
	m.Called(messages)
}

func (m *BrokerMock) AddExchange(exchange *exchange.Exchange) {
	m.Called(exchange)
}

func (m *BrokerMock) AddQueue(queue *queue.Queue) {
	m.Called(queue)
}

func (m *BrokerMock) AddQueueMessageBind(queue *queue.Queue, bindingKeys []string) {
	m.Called(queue, bindingKeys)
}

func (m *BrokerMock) AddQueueConsumer(consumer model.Consumer) {
	m.Called(consumer)
}

func NewBrokerMock() *BrokerMock {
	return new(BrokerMock)
}
