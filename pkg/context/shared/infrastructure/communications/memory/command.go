package memory

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/command"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type (
	commandMapper = map[command.Type]command.Handler
)

type CommandBus struct {
	Handlers commandMapper
}

func (bus *CommandBus) Register(cmd command.Type, handler command.Handler) error {
	_, ok := bus.Handlers[cmd]

	if ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Register",
			What:  fmt.Sprintf("%s already registered", cmd),
			Why: errors.Meta{
				"Command": cmd,
			},
		})
	}

	bus.Handlers[cmd] = handler

	return nil
}

func (bus *CommandBus) Dispatch(cmd command.Command) error {
	handler, ok := bus.Handlers[cmd.Type()]

	if !ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Dispatch",
			What:  "Failure to execute a Command without a Handler",
			Why: errors.Meta{
				"Command": cmd.Type(),
			},
		})
	}

	err := handler.Handle(cmd)

	if err != nil {
		return errors.BubbleUp(err, "Dispatch")
	}

	return nil
}

func NewCommandBus(handlers []command.Handler) (*CommandBus, error) {
	bus := &CommandBus{
		Handlers: make(commandMapper),
	}

	var err error

	for _, handler := range handlers {
		err = bus.Register(handler.SubscribedTo(), handler)

		if err != nil {
			return nil, errors.BubbleUp(err, "NewCommandBus")
		}
	}

	return bus, nil
}
