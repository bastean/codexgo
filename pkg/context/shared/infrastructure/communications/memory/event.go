package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type EventBus struct {
	Consumers events.Mapper
}

func (b *EventBus) Subscribe(key *messages.Key, consumer roles.EventConsumer) error {
	b.Consumers[key.Value()] = append(b.Consumers[key.Value()], consumer)
	return nil
}

func (b *EventBus) Publish(event *messages.Message) error {
	consumers, ok := b.Consumers[event.Key.Value()]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to execute a Event without a Consumer",
			Why: errors.Meta{
				"ID":  event.ID.Value(),
				"Key": event.Key.Value(),
			},
		})
	}

	for _, consumer := range consumers {
		err := consumer.On(event)

		if err != nil {
			return errors.BubbleUp(err)
		}
	}

	return nil
}
