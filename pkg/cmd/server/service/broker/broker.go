package broker

import (
	"os"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/srouter"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/scommunication"
)

var uri = os.Getenv("BROKER_URI")

var Broker smodel.Broker

func Init() error {
	logger.Info("starting rabbitmq")

	rabbitMQ, err := scommunication.NewRabbitMQ(uri, logger.Logger)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	Broker = rabbitMQ

	router := &srouter.Router{
		Name: "codexgo",
	}

	err = Broker.AddRouter(router)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	err = Broker.AddQueue(NotifySendAccountConfirmationQueue)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	err = Broker.AddQueueMessageBind(NotifySendAccountConfirmationQueue, []string{"#.event.#.created.succeeded"})

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	err = Broker.AddQueueConsumer(NotifySendAccountConfirmationQueueConsumer)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	return nil
}

func Close() error {
	logger.Info("closing rabbitmq")

	err := scommunication.CloseRabbitMQ(Broker.(*scommunication.RabbitMQ))

	if err != nil {
		return serror.BubbleUp(err, "Close")
	}

	return nil
}
