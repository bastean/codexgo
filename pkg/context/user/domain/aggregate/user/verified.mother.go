package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

func VerifiedWithValidValue() *Verified {
	value, err := values.New[*Verified](services.Create.Bool())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func VerifiedWithTrueValue() *Verified {
	value, err := values.New[*Verified](true)

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func VerifiedWithFalseValue() *Verified {
	value, err := values.New[*Verified](false)

	if err != nil {
		errors.Panic(err)
	}

	return value
}
