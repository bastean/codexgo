package smodel

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/smessage"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
	"github.com/bastean/codexgo/pkg/context/shared/domain/srouter"
)

type Broker interface {
	AddRouter(router *srouter.Router) error
	AddQueue(queue *squeue.Queue) error
	AddQueueMessageBind(queue *squeue.Queue, bindingKeys []string) error
	AddQueueConsumer(consumer Consumer) error
	PublishMessages(messages []*smessage.Message) error
}
