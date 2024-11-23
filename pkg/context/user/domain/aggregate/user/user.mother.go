package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

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

func RandomPrimitive() *User {
	user, err := FromPrimitive(&Primitive{
		ID:       IDWithValidValue().Value,
		Email:    EmailWithValidValue().Value,
		Username: UsernameWithValidValue().Value,
		Password: PlainPasswordWithValidValue().Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomPrimitive")
	}

	return user
}
