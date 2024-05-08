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

func (broker *BrokerMock) PublishMessages(messages []*message.Message) error {
	args := broker.Called(messages)
	return args.Get(0).(error)
}

func (broker *BrokerMock) AddExchange(exchange *exchange.Exchange) error {
	args := broker.Called(exchange)
	return args.Get(0).(error)
}

func (broker *BrokerMock) AddQueue(queue *queue.Queue) error {
	args := broker.Called(queue)
	return args.Get(0).(error)
}

func (broker *BrokerMock) AddQueueMessageBind(queue *queue.Queue, bindingKeys []string) error {
	args := broker.Called(queue, bindingKeys)
	return args.Get(0).(error)
}

func (broker *BrokerMock) AddQueueConsumer(consumer model.Consumer) error {
	args := broker.Called(consumer)
	return args.Get(0).(error)
}
