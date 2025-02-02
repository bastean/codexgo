package commands

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type (
	Mapper map[messages.Key]roles.CommandHandler
)

func AddCommandMapper(bus roles.CommandBus, mapper Mapper) error {
	var err error

	for key, handler := range mapper {
		err = bus.Register(key, handler)

		if err != nil {
			return errors.BubbleUp(err, "AddCommandMapper")
		}
	}

	return nil
}
