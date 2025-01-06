package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func VerifiedWithValidValue() *Verified {
	value, err := NewVerified(services.Create.Bool())

	if err != nil {
		errors.Panic(err.Error(), "VerifiedWithValidValue")
	}

	return value
}

func VerifiedWithTrueValue() *Verified {
	value, err := NewVerified(true)

	if err != nil {
		errors.Panic(err.Error(), "VerifiedWithTrueValue")
	}

	return value
}

func VerifiedWithFalseValue() *Verified {
	value, err := NewVerified(false)

	if err != nil {
		errors.Panic(err.Error(), "VerifiedWithFalseValue")
	}

	return value
}
