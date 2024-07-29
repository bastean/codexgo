package communications

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/stretchr/testify/mock"
)

type BrokerMock struct {
	mock.Mock
}

func (broker *BrokerMock) PublishMessages(messages []*messages.Message) error {
	broker.Called(messages)
	return nil
}

func (broker *BrokerMock) AddRouter(router *messages.Router) error {
	broker.Called(router)
	return nil
}

func (broker *BrokerMock) AddQueue(queue *messages.Queue) error {
	broker.Called(queue)
	return nil
}

func (broker *BrokerMock) AddQueueMessageBind(queue *messages.Queue, bindingKeys []string) error {
	broker.Called(queue, bindingKeys)
	return nil
}

func (broker *BrokerMock) AddQueueConsumer(consumer messages.Consumer) error {
	broker.Called(consumer)
	return nil
}
