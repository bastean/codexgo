package memory

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type (
	CommandMapper = map[commands.Key]commands.Handler
)

type CommandBus struct {
	Handlers CommandMapper
}

func (bus *CommandBus) Register(key commands.Key, handler commands.Handler) error {
	_, ok := bus.Handlers[key]

	if ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Register",
			What:  fmt.Sprintf("%s already registered", key),
			Why: errors.Meta{
				"Command": key,
			},
		})
	}

	bus.Handlers[key] = handler

	return nil
}

func (bus *CommandBus) Dispatch(command *commands.Command) error {
	handler, ok := bus.Handlers[command.Key]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Dispatch",
			What:  "Failure to execute a Command without a Handler",
			Why: errors.Meta{
				"Command": command.Key,
			},
		})
	}

	err := handler.Handle(command)

	if err != nil {
		return errors.BubbleUp(err, "Dispatch")
	}

	return nil
}

func NewCommandBus(mapper CommandMapper) (*CommandBus, error) {
	bus := &CommandBus{
		Handlers: make(CommandMapper),
	}

	var err error

	for key, handler := range mapper {
		err = bus.Register(key, handler)

		if err != nil {
			return nil, errors.BubbleUp(err, "NewCommandBus")
		}
	}

	return bus, nil
}
