package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type (
	EventMapper = map[messages.Key][]roles.EventConsumer
)

type EventBus struct {
	Consumers EventMapper
}

func (b *EventBus) Subscribe(key messages.Key, consumer roles.EventConsumer) error {
	b.Consumers[key] = append(b.Consumers[key], consumer)
	return nil
}

func (b *EventBus) Publish(event *messages.Message) error {
	consumers, ok := b.Consumers[event.Key]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Publish",
			What:  "Failure to execute a Event without a Consumer",
			Why: errors.Meta{
				"Event": event.Key,
			},
		})
	}

	for _, consumer := range consumers {
		err := consumer.On(event)

		if err != nil {
			return errors.BubbleUp(err, "Publish")
		}
	}

	return nil
}

func NewEventBus(mapper EventMapper) (*EventBus, error) {
	bus := &EventBus{
		Consumers: make(EventMapper, len(mapper)),
	}

	var err error

	for key, consumers := range mapper {
		for _, consumer := range consumers {
			err = bus.Subscribe(key, consumer)

			if err != nil {
				return nil, errors.BubbleUp(err, "NewEventBus")
			}
		}
	}

	return bus, nil
}
