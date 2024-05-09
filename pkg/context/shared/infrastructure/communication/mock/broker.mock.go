package communicationMock

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/router"
	"github.com/stretchr/testify/mock"
)

type BrokerMock struct {
	mock.Mock
}

func (broker *BrokerMock) PublishMessages(messages []*message.Message) error {
	broker.Called(messages)
	return nil
}

func (broker *BrokerMock) AddRouter(router *router.Router) error {
	broker.Called(router)
	return nil
}

func (broker *BrokerMock) AddQueue(queue *queue.Queue) error {
	broker.Called(queue)
	return nil
}

func (broker *BrokerMock) AddQueueMessageBind(queue *queue.Queue, bindingKeys []string) error {
	broker.Called(queue, bindingKeys)
	return nil
}

func (broker *BrokerMock) AddQueueConsumer(consumer model.Consumer) error {
	broker.Called(consumer)
	return nil
}
