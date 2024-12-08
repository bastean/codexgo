package rabbitmq

import (
	"context"
	"reflect"

	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/loggers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
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

type (
	RabbitMQ = rabbitmq.RabbitMQ
	Events   = rabbitmq.Events
)

func Open(uri string, exchange string, queues rabbitmq.Queues, mapper rabbitmq.Events, logger loggers.Logger, consumeCycle context.Context) (*rabbitmq.RabbitMQ, error) {
	rmq, err := rabbitmq.Open(
		uri,
		exchange,
		queues,
		logger,
		consumeCycle,
	)

	if err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	for key, consumers := range mapper {
		for _, consumer := range consumers {
			rmq.Subscribe(key, consumer)
		}
	}

	return rmq, nil
}
