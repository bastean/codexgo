package memory

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
)

type (
	QueryMapper = map[messages.Key]queries.Handler
)

type QueryBus struct {
	Handlers QueryMapper
}

func (b *QueryBus) Register(key messages.Key, handler queries.Handler) error {
	_, ok := b.Handlers[key]

	if ok {
		return errors.New[errors.Internal](&errors.Bubble{
			Where: "Register",
			What:  fmt.Sprintf("%s already registered", key),
			Why: errors.Meta{
				"Query": key,
			},
		})
	}

	b.Handlers[key] = handler

	return nil
}

func (b *QueryBus) Ask(query *messages.Message) (*messages.Message, error) {
	handler, ok := b.Handlers[query.Key]

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
		Handlers: make(QueryMapper, len(mapper)),
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
