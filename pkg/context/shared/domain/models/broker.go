package models

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
	"github.com/bastean/codexgo/pkg/context/shared/domain/routers"
)

type Broker interface {
	AddRouter(router *routers.Router) error
	AddQueue(queue *queues.Queue) error
	AddQueueMessageBind(queue *queues.Queue, bindingKeys []string) error
	AddQueueConsumer(consumer Consumer) error
	PublishMessages(messages []*messages.Message) error
}
