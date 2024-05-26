package broker

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
)

var NotifySendAccountConfirmationQueue = &queues.Queue{
	Name: queues.NewQueueName(&queues.QueueName{
		Module: "notify",
		Action: "send account confirmation",
		Event:  "created.succeeded",
	}),
}
