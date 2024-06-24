package rabbitmq

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
)

var UserSendConfirmationQueue = &queues.Queue{
	Name: queues.NewQueueName(&queues.QueueNameComponents{
		Service: "user",
		Entity:  "user",
		Action:  "send confirmation",
		Event:   "created succeeded",
	}),
	Bindings: []string{Binding.Event.CreatedSucceeded},
}

var Queues = []*queues.Queue{
	UserSendConfirmationQueue,
}
