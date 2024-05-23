package broker

import (
	"os"

	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/cmd/server/service/notify"
	"github.com/bastean/codexgo/pkg/context/notify/application/send"
	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/squeue"
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

	notifySendAccountConfirmationQueueName := squeue.NewQueueName(&squeue.QueueName{
		Module: "notify",
		Action: "send account confirmation",
		Event:  "created.succeeded",
	})

	notifySendAccountConfirmationQueue := &squeue.Queue{
		Name: notifySendAccountConfirmationQueueName,
	}

	notifySendAccountConfirmationQueueConsumer := &send.CreatedSucceededEventConsumer{
		UseCase: notify.SendAccountConfirmationMail,
		Queues:  []*squeue.Queue{notifySendAccountConfirmationQueue},
	}

	err = Broker.AddRouter(router)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	err = Broker.AddQueue(notifySendAccountConfirmationQueue)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	err = Broker.AddQueueMessageBind(notifySendAccountConfirmationQueue, []string{"#.event.#.created.succeeded"})

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	err = Broker.AddQueueConsumer(notifySendAccountConfirmationQueueConsumer)

	if err != nil {
		return serror.BubbleUp(err, "Init")
	}

	return nil
}

func Close() error {
	err := scommunication.CloseRabbitMQ(Broker.(*scommunication.RabbitMQ))

	if err != nil {
		return serror.BubbleUp(err, "Close")
	}

	logger.Info("rabbitmq closed")

	return nil
}
