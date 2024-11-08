package memory

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
)

type (
	queryMapper = map[queries.Type]queries.Handler
)

type QueryBus struct {
	Handlers queryMapper
}

func (bus *QueryBus) Register(ask queries.Type, handler queries.Handler) error {
	_, ok := bus.Handlers[ask]

	if ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Register",
			What:  fmt.Sprintf("%s already registered", ask),
			Why: errors.Meta{
				"Query": ask,
			},
		})
	}

	bus.Handlers[ask] = handler

	return nil
}

func (bus *QueryBus) Ask(ask queries.Query) (queries.Response, error) {
	handler, ok := bus.Handlers[ask.Type()]

	if !ok {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Ask",
			What:  "Failure to execute a Query without a Handler",
			Why: errors.Meta{
				"Query": ask.Type(),
			},
		})
	}

	response, err := handler.Handle(ask)

	if err != nil {
		return nil, errors.BubbleUp(err, "Ask")
	}

	return response, nil
}

func NewQueryBus(handlers []queries.Handler) (*QueryBus, error) {
	bus := &QueryBus{
		Handlers: make(queryMapper),
	}

	var err error

	for _, handler := range handlers {
		err = bus.Register(handler.SubscribedTo(), handler)

		if err != nil {
			return nil, errors.BubbleUp(err, "NewQueryBus")
		}
	}

	return bus, nil
}
