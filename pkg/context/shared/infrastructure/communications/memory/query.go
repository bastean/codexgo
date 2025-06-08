package memory

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type QueryBus struct {
	Handlers queries.Mapper
}

func (b *QueryBus) Register(key *messages.Key, handler roles.QueryHandler) error {
	_, ok := b.Handlers[key.Value()]

	if ok {
		return errors.New[errors.Internal](&errors.Bubble{
			What: "Already registered",
			Why: errors.Meta{
				"Key": key.Value(),
			},
		})
	}

	b.Handlers[key.Value()] = handler

	return nil
}

func (b *QueryBus) Ask(query *messages.Message) (*messages.Message, error) {
	handler, ok := b.Handlers[query.Key.Value()]

	if !ok {
		return nil, errors.New[errors.Internal](&errors.Bubble{
			What: "Failure to execute a Query without a Handler",
			Why: errors.Meta{
				"ID":  query.ID.Value(),
				"Key": query.Key.Value(),
			},
		})
	}

	response, err := handler.Handle(query)

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	return response, nil
}
