package communication

import (
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/memory"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/notification"
	"github.com/bastean/codexgo/v4/internal/pkg/service/record/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
)

var Service = &struct {
	RabbitMQ, InMemory string
}{
	RabbitMQ: log.Service("RabbitMQ"),
	InMemory: log.Service("InMemory"),
}

var (
	err      error
	RabbitMQ *rabbitmq.RabbitMQ
	QueryBus *memory.QueryBus
)

func Up() error {
	log.EstablishingConnectionWith(Service.RabbitMQ)

	RabbitMQ, err = rabbitmq.Open(
		env.BrokerRabbitMQURI,
		env.BrokerRabbitMQName,
		rabbitmq.Queues,
		rabbitmq.Events{
			user.CreatedSucceededKey: []events.Consumer{
				notification.Confirmation,
			},
		},
		log.Log,
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
