package model

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/exchange"
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
)

type Broker interface {
	AddExchange(exchange *exchange.Exchange) error
	AddQueue(queue *queue.Queue) error
	AddQueueMessageBind(queue *queue.Queue, bindingKeys []string) error
	AddQueueConsumer(consumer Consumer) error
	PublishMessages(messages []*message.Message) error
}
