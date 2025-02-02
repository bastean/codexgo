package queries

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type (
	Mapper map[messages.Key]roles.QueryHandler
)

func AddQueryMapper(bus roles.QueryBus, mapper Mapper) error {
	var err error

	for key, handler := range mapper {
		err = bus.Register(key, handler)

		if err != nil {
			return errors.BubbleUp(err, "AddQueryMapper")
		}
	}

	return nil
}
