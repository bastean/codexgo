package communication

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/memory"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
	"github.com/bastean/codexgo/v4/internal/pkg/service/record/log"
)

var Service = &struct {
	RabbitMQ, InMemory string
}{
	RabbitMQ: log.Service("RabbitMQ"),
	InMemory: log.Service("InMemory"),
}

var (
	err        error
	RabbitMQ   *rabbitmq.RabbitMQ
	CommandBus *memory.CommandBus
)

func Up() error {
	log.EstablishingConnectionWith(Service.RabbitMQ)

	RabbitMQ, err = rabbitmq.Open(
		env.BrokerRabbitMQURI,
		log.Log,
		rabbitmq.Exchange(env.BrokerRabbitMQName),
		rabbitmq.Queues{
			user.QueueSendConfirmation,
		},
		rabbitmq.Consumers{
			user.Created,
		},
	)

	if err != nil {
		log.ConnectionFailedWith(Service.RabbitMQ)
		return errors.BubbleUp(err, "Up")
	}

	log.ConnectionEstablishedWith(Service.RabbitMQ)

	return nil
}

func Down() error {
	log.ClosingConnectionWith(Service.RabbitMQ)

	if err = rabbitmq.Close(RabbitMQ); err != nil {
		log.DisconnectionFailedWith(Service.RabbitMQ)
		return errors.BubbleUp(err, "Down")
	}

	log.ConnectionClosedWith(Service.RabbitMQ)

	return nil
}
