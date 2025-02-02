package communication

import (
	"context"

	"github.com/bastean/codexgo/v4/internal/pkg/adapter/command"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/event"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/internal/pkg/adapter/query"
	"github.com/bastean/codexgo/v4/internal/pkg/service/env"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/rabbitmq"
)

var Service = &struct {
	RabbitMQEventBus, MemoryEventBus, MemoryCommandBus, MemoryQueryBus string
}{
	RabbitMQEventBus: log.Service("RabbitMQEventBus"),
	MemoryEventBus:   log.Service("MemoryEventBus"),
	MemoryCommandBus: log.Service("MemoryCommandBus"),
	MemoryQueryBus:   log.Service("MemoryQueryBus"),
}

var (
	err error
)

func Up() error {
	switch {
	case env.HasRabbitMQ():
		log.EstablishingConnectionWith(Service.RabbitMQEventBus)

		event.Bus, err = rabbitmq.Open(
			env.BrokerRabbitMQURI,
			env.BrokerRabbitMQName,
			log.Log,
			context.Background(),
		)

		if err != nil {
			log.ConnectionFailedWith(Service.RabbitMQEventBus)
			return errors.BubbleUp(err, "Up")
		}

		err = rabbitmq.AddQueueMapper(event.Bus.(*rabbitmq.RabbitMQ), event.RabbitMQueueMapper)

		if err != nil {
			log.CannotBeStarted(Service.RabbitMQEventBus)
			return errors.BubbleUp(err, "Up")
		}

		log.ConnectionEstablishedWith(Service.RabbitMQEventBus)
	default:
		log.Starting(Service.MemoryEventBus)

		event.Bus = &memory.EventBus{
			Consumers: events.Mapper{},
		}

		log.Started(Service.MemoryEventBus)
	}

	switch {
	default:
		log.Starting(Service.MemoryCommandBus)

		command.Bus = &memory.CommandBus{
			Handlers: commands.Mapper{},
		}

		log.Started(Service.MemoryCommandBus)
	}

	switch {
	default:
		log.Starting(Service.MemoryQueryBus)

		query.Bus = &memory.QueryBus{
			Handlers: queries.Mapper{},
		}

		log.Started(Service.MemoryQueryBus)
	}

	return nil
}

func Down() error {
	switch {
	case env.HasRabbitMQ():
		log.ClosingConnectionWith(Service.RabbitMQEventBus)

		if err = rabbitmq.Close(event.Bus.(*rabbitmq.RabbitMQ)); err != nil {
			log.DisconnectionFailedWith(Service.RabbitMQEventBus)
			return errors.BubbleUp(err, "Down")
		}

		log.ConnectionClosedWith(Service.RabbitMQEventBus)
	}

	return nil
}
