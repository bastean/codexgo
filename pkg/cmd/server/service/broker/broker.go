package broker

import (
	"os"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/routers"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communications"
)

var uri = os.Getenv("BROKER_URI")

var Broker models.Broker

func Init() error {
	logger.Info("starting rabbitmq")

	rabbitMQ, err := communications.NewRabbitMQ(uri, logger.Logger)

	if err != nil {
		return errors.BubbleUp(err, "Init")
	}

	Broker = rabbitMQ

	router := &routers.Router{
		Name: "codexgo",
	}

	err = Broker.AddRouter(router)

	if err != nil {
		return errors.BubbleUp(err, "Init")
	}

	err = Broker.AddQueue(NotifySendAccountConfirmationQueue)

	if err != nil {
		return errors.BubbleUp(err, "Init")
	}

	err = Broker.AddQueueMessageBind(NotifySendAccountConfirmationQueue, []string{"#.event.#.created.succeeded"})

	if err != nil {
		return errors.BubbleUp(err, "Init")
	}

	err = Broker.AddQueueConsumer(NotifySendAccountConfirmationQueueConsumer)

	if err != nil {
		return errors.BubbleUp(err, "Init")
	}

	return nil
}

func Close() error {
	logger.Info("closing rabbitmq")

	err := communications.CloseRabbitMQ(Broker.(*communications.RabbitMQ))

	if err != nil {
		return errors.BubbleUp(err, "Close")
	}

	return nil
}
