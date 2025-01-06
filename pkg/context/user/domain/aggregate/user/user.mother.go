package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func Random() *User {
	user, err := New(&Primitive{
		ID:       IDWithValidValue().Value,
		Email:    EmailWithValidValue().Value,
		Username: UsernameWithValidValue().Value,
		Password: PlainPasswordWithValidValue().Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "Random")
	}

	return user
}

func RandomPrimitive() *User {
	user, err := FromPrimitive(&Primitive{
		Created:  aggregates.TimeWithValidValue().Value,
		Updated:  aggregates.TimeWithValidValue().Value,
		ID:       IDWithValidValue().Value,
		Email:    EmailWithValidValue().Value,
		Username: UsernameWithValidValue().Value,
		Password: PlainPasswordWithValidValue().Value,
		Verified: VerifiedWithValidValue().Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomPrimitive")
	}

	return user
}

func RandomRaw() *User {
	user := RandomPrimitive()

	user.Created = nil
	user.Updated = nil

	return user
}
