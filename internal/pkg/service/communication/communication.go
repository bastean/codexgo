package communication

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/event"
	"github.com/bastean/codexgo/v4/internal/pkg/service/communication/rabbitmq"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/internal/pkg/service/errors"
	"github.com/bastean/codexgo/v4/internal/pkg/service/module/notification"
	"github.com/bastean/codexgo/v4/internal/pkg/service/record/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
)

var Service = &struct {
	RabbitMQ, EventBus, CommandBus, QueryBus string
}{
	RabbitMQ:   log.Service("RabbitMQ"),
	EventBus:   log.Service("EventBus"),
	CommandBus: log.Service("CommandBus"),
	QueryBus:   log.Service("QueryBus"),
}

var (
	err error
	Bus event.Bus
)

func Up() error {
	switch {
	case env.HasBroker():
		log.EstablishingConnectionWith(Service.RabbitMQ)

		Bus, err = rabbitmq.Open(
			env.BrokerRabbitMQURI,
			env.BrokerRabbitMQName,
			rabbitmq.Queues,
			rabbitmq.Events{
				events.UserCreatedSucceededKey: []event.Consumer{
					notification.Confirmation,
				},
			},
			log.Log,
			context.Background(),
		)

		if err != nil {
			log.ConnectionFailedWith(Service.RabbitMQ)
			return errors.BubbleUp(err, "Up")
		}

		log.ConnectionEstablishedWith(Service.RabbitMQ)
	default:
		log.Starting(Service.EventBus)

		Bus, err = event.NewBus(event.Mapper{
			events.UserCreatedSucceededKey: []event.Consumer{
				notification.Confirmation,
			},
		})

		if err != nil {
			log.CannotBeStarted(Service.EventBus)
			return errors.BubbleUp(err, "Up")
		}

		log.Started(Service.EventBus)
	}

	return nil
}

func Down() error {
	switch {
	case env.HasBroker():
		log.ClosingConnectionWith(Service.RabbitMQ)

		if err = rabbitmq.Close(Bus.(*rabbitmq.RabbitMQ)); err != nil {
			log.DisconnectionFailedWith(Service.RabbitMQ)
			return errors.BubbleUp(err, "Down")
		}

		log.ConnectionClosedWith(Service.RabbitMQ)
	}

	return nil
}
