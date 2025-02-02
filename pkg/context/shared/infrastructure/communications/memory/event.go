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
