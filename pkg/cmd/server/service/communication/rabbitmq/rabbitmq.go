package rabbitmq

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications"
)

func New(uri string, logger models.Logger, exchange *messages.Router, queues []*messages.Queue, consumers []messages.Consumer) (messages.Broker, error) {
	rabbitMQ, err := communications.NewRabbitMQ(uri, logger)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	err = rabbitMQ.AddRouter(exchange)

	if err != nil {
		return nil, errors.BubbleUp(err, "New")
	}

	for _, queue := range queues {

		err = rabbitMQ.AddQueue(queue)

		if err != nil {
			return nil, errors.BubbleUp(err, "New")
		}

		err = rabbitMQ.AddQueueMessageBind(queue, queue.Bindings)

		if err != nil {
			return nil, errors.BubbleUp(err, "New")
		}
	}

	for _, consumer := range consumers {
		err = rabbitMQ.AddQueueConsumer(consumer)

		if err != nil {
			return nil, errors.BubbleUp(err, "New")
		}
	}

	return rabbitMQ, nil
}

func Close(rabbitMQ messages.Broker) error {
	err := communications.CloseRabbitMQ(rabbitMQ.(*communications.RabbitMQ))

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
