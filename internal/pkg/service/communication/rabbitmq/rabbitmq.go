package rabbitmq

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/loggers"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications/rabbitmq"
)

func New(uri string, logger loggers.Logger, exchange *messages.Router, queues []*messages.Queue, consumers []messages.Consumer) (messages.Broker, error) {
	rabbitMQ, err := rabbitmq.New(uri, logger)

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

func Close(connection messages.Broker) error {
	err := rabbitmq.Close(connection.(*rabbitmq.RabbitMQ))

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
