package memory

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
)

type (
	QueryMapper = map[queries.Key]queries.Handler
)

type QueryBus struct {
	Handlers QueryMapper
}

func (bus *QueryBus) Register(key queries.Key, handler queries.Handler) error {
	_, ok := bus.Handlers[key]

	if ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Register",
			What:  fmt.Sprintf("%s already registered", key),
			Why: errors.Meta{
				"Query": key,
			},
		})
	}

	bus.Handlers[key] = handler

	return nil
}

func (bus *QueryBus) Ask(query *queries.Query) (*queries.Response, error) {
	handler, ok := bus.Handlers[query.Key]

	if !ok {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			Where: "Ask",
			What:  "Failure to execute a Query without a Handler",
			Why: errors.Meta{
				"Query": query.Key,
			},
		})
	}

	response, err := handler.Handle(query)

	if err != nil {
		return nil, errors.BubbleUp(err, "Ask")
	}

	return response, nil
}

func NewQueryBus(mapper QueryMapper) (*QueryBus, error) {
	bus := &QueryBus{
		Handlers: make(QueryMapper),
	}

	var err error

	for key, handler := range mapper {
		err = bus.Register(key, handler)

		if err != nil {
			return nil, errors.BubbleUp(err, "NewQueryBus")
		}
	}

	return bus, nil
}
