package scommunication

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/srouter"
	"github.com/stretchr/testify/mock"
)

type BrokerMock struct {
	mock.Mock
}

func (broker *BrokerMock) PublishMessages(messages []*smessage.Message) error {
	broker.Called(messages)
	return nil
}

func (broker *BrokerMock) AddRouter(router *srouter.Router) error {
	broker.Called(router)
	return nil
}

func (broker *BrokerMock) AddQueue(queue *squeue.Queue) error {
	broker.Called(queue)
	return nil
}

func (broker *BrokerMock) AddQueueMessageBind(queue *squeue.Queue, bindingKeys []string) error {
	broker.Called(queue, bindingKeys)
	return nil
}

func (broker *BrokerMock) AddQueueConsumer(consumer smodel.Consumer) error {
	broker.Called(consumer)
	return nil
}
