package communications

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
	"github.com/bastean/codexgo/pkg/context/shared/domain/routers"
	"github.com/stretchr/testify/mock"
)

type BrokerMock struct {
	mock.Mock
}

func (broker *BrokerMock) PublishMessages(messages []*messages.Message) error {
	broker.Called(messages)
	return nil
}

func (broker *BrokerMock) AddRouter(router *routers.Router) error {
	broker.Called(router)
	return nil
}

func (broker *BrokerMock) AddQueue(queue *queues.Queue) error {
	broker.Called(queue)
	return nil
}

func (broker *BrokerMock) AddQueueMessageBind(queue *queues.Queue, bindingKeys []string) error {
	broker.Called(queue, bindingKeys)
	return nil
}

func (broker *BrokerMock) AddQueueConsumer(consumer models.Consumer) error {
	broker.Called(consumer)
	return nil
}
