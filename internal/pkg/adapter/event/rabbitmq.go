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
}
