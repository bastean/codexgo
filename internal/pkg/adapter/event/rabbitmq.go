package event

import (
	"reflect"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
)

var RabbitMQueueMapper = rabbitmq.QueueMapper{
	events.UserCreatedSucceededKey: &rabbitmq.Recipient{
		Name: messages.NewRecipient(&messages.RecipientComponents{
			Service: "user",
			Entity:  "user",
			Action:  "send confirmation",
			Event:   "created",
			Status:  "succeeded",
		}),
		BindingKey: events.UserCreatedSucceededKey,
		Attributes: reflect.TypeOf(new(events.UserCreatedSucceededAttributes)),
	},
	events.UserResetQueuedKey: &rabbitmq.Recipient{
		Name: messages.NewRecipient(&messages.RecipientComponents{
			Service: "user",
			Entity:  "user",
			Action:  "send reset",
			Event:   "reset",
			Status:  "queued",
		}),
		BindingKey: events.UserResetQueuedKey,
		Attributes: reflect.TypeOf(new(events.UserResetQueuedAttributes)),
	},
}
