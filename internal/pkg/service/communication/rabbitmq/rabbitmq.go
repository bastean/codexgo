package rabbitmq

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/loggers"
	"github.com/bastean/codexgo/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications/rabbitmq"
)

func Open(uri string, logger loggers.Logger, exchange *messages.Router, queues []*messages.Queue, consumers []messages.Consumer) (messages.Broker, error) {
	session, err := rabbitmq.Open(uri, logger)

	if err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	err = session.AddRouter(exchange)

	if err != nil {
		return nil, errors.BubbleUp(err, "Open")
	}

	for _, queue := range queues {

		err = session.AddQueue(queue)

		if err != nil {
			return nil, errors.BubbleUp(err, "Open")
		}

		err = session.AddQueueMessageBind(queue, queue.Bindings)

		if err != nil {
			return nil, errors.BubbleUp(err, "Open")
		}
	}

	for _, consumer := range consumers {
		err = session.AddQueueConsumer(consumer)

		if err != nil {
			return nil, errors.BubbleUp(err, "Open")
		}
	}

	return session, nil
}

func Close(session messages.Broker) error {
	err := rabbitmq.Close(session.(*rabbitmq.RabbitMQ))

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
