package rabbitmq

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queues"
	"github.com/bastean/codexgo/pkg/context/shared/domain/routers"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications"
)

func New(uri string, logger models.Logger, exchange *routers.Router, queues []*queues.Queue, consumers []models.Consumer) (models.Broker, error) {
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

		err = rabbitMQ.AddQueueMessageBind(queue, queue.BindingKeys)

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

func Close(rabbitMQ models.Broker) error {
	err := communications.CloseRabbitMQ(rabbitMQ.(*communications.RabbitMQ))

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
