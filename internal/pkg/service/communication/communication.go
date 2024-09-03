package communication

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/logger/log"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/user"
)

var Service = &struct {
	RabbitMQ string
}{
	RabbitMQ: log.Service("RabbitMQ"),
}

var (
	err      error
	RabbitMQ *rabbitmq.RabbitMQ
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
