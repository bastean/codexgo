package components

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func TypeWithValidValue() *Type {
	value, err := NewType(services.Create.RandomString([]string{"event", "command", "query", "response"}))

	if err != nil {
		errors.Panic(err.Error(), "TypeWithValidValue")
	}

	return value
}

func TypeWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewType(value)

	return value, err
}
