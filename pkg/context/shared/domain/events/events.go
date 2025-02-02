package events

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type (
	Mapper map[messages.Key][]roles.EventConsumer
)

func AddEventMapper(bus roles.EventBus, mapper Mapper) error {
	var err error

	for key, consumers := range mapper {
		for _, consumer := range consumers {
			err = bus.Subscribe(key, consumer)

			if err != nil {
				return errors.BubbleUp(err, "AddEventMapper")
			}
		}
	}

	return nil
}
