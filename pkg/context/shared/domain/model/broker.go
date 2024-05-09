package model

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/router"
)

type Broker interface {
	AddRouter(router *router.Router) error
	AddQueue(queue *queue.Queue) error
	AddQueueMessageBind(queue *queue.Queue, bindingKeys []string) error
	AddQueueConsumer(consumer Consumer) error
	PublishMessages(messages []*message.Message) error
}
