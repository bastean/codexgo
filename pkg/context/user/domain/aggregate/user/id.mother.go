package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

func IDWithValidValue() *ID {
	value, err := values.New[*ID](services.Create.UUID())

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func IDWithInvalidValue() (string, error) {
	value := "x"

	_, err := values.New[*ID](value)

	return value, err
}
