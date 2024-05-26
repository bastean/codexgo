package models

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
)

type Consumer interface {
	SubscribedTo() []*queues.Queue
	On(message *messages.Message) error
}
