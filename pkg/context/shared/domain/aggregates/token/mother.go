package token

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/attempt"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mother"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type m struct {
	*mother.Mother
}

func (m *m) TokenNew(id string) *Token {
	token, err := New(id)

	if err != nil {
		errors.Panic(err)
	}

	return token
}

func (m *m) TokenCopy(token *Token) *Token {
	copy, err := FromPrimitive(token.ToPrimitive())

	if err != nil {
		errors.Panic(err)
	}

	return copy
}

func (m *m) TokenValid() *Token {
	token, err := New(values.Mother().IDValid().Value())

	if err != nil {
		errors.Panic(err)
	}

	return token
}

func (m *m) TokenValidFromPrimitive() *Token {
	token, err := FromPrimitive(&Primitive{
		ID:      values.Mother().IDValid().ToPrimitive(),
		Attempt: attempt.Mother().AttemptValid().ToPrimitive(),
	})

	if err != nil {
		errors.Panic(err)
	}

	return token
}

var Mother = mother.New[m]
