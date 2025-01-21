package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

func RandomPrimitive() *User {
	user, err := FromPrimitive(&Primitive{
		Created:  aggregates.TimeWithValidValue().Value,
		Updated:  aggregates.TimeWithValidValue().Value,
		Verify:   IDWithValidValue().Value,
		Reset:    IDWithValidValue().Value,
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
	user, err := FromRaw(&Primitive{
		ID:       IDWithValidValue().Value,
		Email:    EmailWithValidValue().Value,
		Username: UsernameWithValidValue().Value,
		Password: PlainPasswordWithValidValue().Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomRaw")
	}

	return user
}

func Random() *User {
	user := RandomPrimitive()

	user.Created = nil
	user.Updated = nil
	user.Reset = nil

	return user
}
