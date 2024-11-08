package rabbitmq

import (
	"reflect"

	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/loggers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
)

type (
	RabbitMQ = rabbitmq.RabbitMQ
	Events   = rabbitmq.Events
)

var (
	Close = rabbitmq.Close
)

var Queues = rabbitmq.Queues{
	user.CreatedSucceededKey: &rabbitmq.Recipient{
		Name: messages.NewRecipient(&messages.RecipientComponents{
			Service: "user",
			Entity:  "user",
			Action:  "send confirmation",
			Event:   "created",
			Status:  "succeeded",
		}),
		BindingKey: user.CreatedSucceededKey,
		Attributes: reflect.TypeOf(new(user.CreatedSucceededAttributes)),
	},
}

func Open(uri string, exchange string, queues rabbitmq.Queues, mapper rabbitmq.Events, logger loggers.Logger) (*rabbitmq.RabbitMQ, error) {
	rmq, err := rabbitmq.Open(
		uri,
		exchange,
		queues,
		logger,
	)

	if err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	for key, consumers := range mapper {
		for _, consumer := range consumers {
			go rmq.Subscribe(key, consumer)
		}
	}

	return rmq, nil
}
