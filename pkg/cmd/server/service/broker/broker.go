package broker

import (
	"github.com/bastean/codexgo/pkg/cmd/server/service/logger"
	"github.com/bastean/codexgo/pkg/cmd/server/service/notify"
	"github.com/bastean/codexgo/pkg/context/notify/application/sendMail"
	"github.com/bastean/codexgo/pkg/context/shared/domain/exchange"
	"github.com/bastean/codexgo/pkg/context/shared/domain/queue"
	"github.com/bastean/codexgo/pkg/context/shared/infrastructure/communication"
)

var Broker = communication.NewRabbitMQ()

var Exchange = exchange.NewExchange("codexgo")

var NotifySendAccountConfirmationQueueName = queue.NewQueueName(&queue.QueueName{Module: "notify", Action: "send account confirmation", Event: "registered.succeeded"})

var NotifySendAccountConfirmationQueue = queue.NewQueue(NotifySendAccountConfirmationQueueName)

var NotifySendAccountConfirmationQueueConsumer = sendMail.NewRegisteredSucceededEventConsumer(notify.NotifySendMail, []*queue.Queue{NotifySendAccountConfirmationQueue})

func init() {
	logger.Logger.Info("starting rabbitmq")

	Broker.AddExchange(Exchange)

	Broker.AddQueue(NotifySendAccountConfirmationQueue)

	Broker.AddQueueMessageBind(NotifySendAccountConfirmationQueue, []string{"#.event.#.registered.succeeded"})

	Broker.AddQueueConsumer(NotifySendAccountConfirmationQueueConsumer)
}
