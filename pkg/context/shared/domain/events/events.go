package events

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type (
	Mapper map[string][]roles.EventConsumer
)

func AddEventMapper(bus roles.EventBus, mapper Mapper) error {
	var (
		err error
		key *messages.Key
	)

	for rawKey, consumers := range mapper {
		key, _ = values.New[*messages.Key](rawKey)

		for _, consumer := range consumers {
			err = bus.Subscribe(key, consumer)

			if err != nil {
				return errors.BubbleUp(err)
			}
		}
	}

	return nil
}
