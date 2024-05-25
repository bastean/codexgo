package broker

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
)

var NotifySendAccountConfirmationQueue = &squeue.Queue{
	Name: squeue.NewQueueName(&squeue.QueueName{
		Module: "notify",
		Action: "send account confirmation",
		Event:  "created.succeeded",
	}),
}
