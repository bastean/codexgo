package model

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/message"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
)

type Consumer interface {
	SubscribedTo() []*queue.Queue
	On(message *message.Message) error
}
