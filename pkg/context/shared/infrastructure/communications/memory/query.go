package memory

import (
	"fmt"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type QueryBus struct {
	Handlers queries.Mapper
}

func (b *QueryBus) Register(key messages.Key, handler roles.QueryHandler) error {
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
