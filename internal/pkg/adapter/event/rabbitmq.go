package event

import (
	"reflect"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
)

var UserCreatedSucceededRecipient, _ = values.New[*messages.Recipient](messages.ParseRecipient(&messages.RecipientComponents{
	Service: "user",
	Entity:  "user",
	Trigger: "send_confirmation",
	Action:  "created",
	Status:  "succeeded",
}))

var UserResetQueuedRecipient, _ = values.New[*messages.Recipient](messages.ParseRecipient(&messages.RecipientComponents{
	Service: "user",
	Entity:  "user",
	Trigger: "send_reset",
	Action:  "reset",
	Status:  "queued",
}))

var RabbitMQueueMapper = rabbitmq.Mapper{
	events.UserCreatedSucceededKey: []*rabbitmq.Queue{
		{
			Name:       UserCreatedSucceededRecipient,
			BindingKey: events.UserCreatedSucceededKey.Value(),
			Attributes: reflect.TypeOf(new(events.UserCreatedSucceededAttributes)),
		},
	},
	events.UserResetQueuedKey: []*rabbitmq.Queue{
		{
			Name:       UserResetQueuedRecipient,
			BindingKey: events.UserResetQueuedKey.Value(),
			Attributes: reflect.TypeOf(new(events.UserResetQueuedAttributes)),
		},
	},
}
