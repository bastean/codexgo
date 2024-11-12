package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func IDWithValidValue() *ID {
	value, err := NewID(services.Create.UUID())

	if err != nil {
		errors.Panic(err.Error(), "IDWithValidValue")
	}

	return value
}

func IDWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewID(value)

	return value, err
}
