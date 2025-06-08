package commands

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type (
	Mapper map[string]roles.CommandHandler
)

func AddCommandMapper(bus roles.CommandBus, mapper Mapper) error {
	var (
		err error
		key *messages.Key
	)

	for rawKey, handler := range mapper {
		key, _ = values.New[*messages.Key](rawKey)

		err = bus.Register(key, handler)

		if err != nil {
			return errors.BubbleUp(err)
		}
	}

	return nil
}
