package token

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/attempt"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

const (
	Limit = 3
	Every = int(time.Minute * 3)
	Next  = int(time.Hour)
)

type Token struct {
	*values.ID
	Attempt *attempt.Attempt
}

type Primitive struct {
	ID      *values.StringPrimitive
	Attempt *attempt.Primitive
}

func (t *Token) ToPrimitive() *Primitive {
	return &Primitive{
		ID:      t.ID.ToPrimitive(),
		Attempt: t.Attempt.ToPrimitive(),
	}
}

func FromPrimitive(primitive *Primitive, isOptional ...bool) (*Token, error) {
	if primitive == nil {
		switch {
		case len(isOptional) == 1:
			return nil, nil
		default:
			return nil, errors.New[errors.Internal](&errors.Bubble{
				What: "Token value is required",
			})
		}
	}

	id, errID := values.FromPrimitive[*values.ID](primitive.ID)
	attempt, errAttempt := attempt.FromPrimitive(primitive.Attempt)

	if err := errors.Join(errID, errAttempt); err != nil {
		return nil, errors.BubbleUp(err)
	}

	return &Token{
		ID:      id,
		Attempt: attempt,
	}, nil
}

func New(token string) (*Token, error) {
	id, errID := values.New[*values.ID](token)

	attempt, errAttempt := attempt.New(&attempt.Required{
		Limit: Limit,
		Every: Every,
		Next:  Next,
	})

	if err := errors.Join(errID, errAttempt); err != nil {
		return nil, errors.BubbleUp(err)
	}

	aggregate := &Token{
		ID:      id,
		Attempt: attempt,
	}

	err := aggregate.Attempt.Increase()

	if err != nil {
		return nil, errors.BubbleUp(err)
	}

	return aggregate, nil
}
