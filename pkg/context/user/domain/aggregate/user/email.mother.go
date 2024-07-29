package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func EmailWithValidValue() *Email {
	value, err := NewEmail(services.Create.Email())

	if err != nil {
		errors.Panic(err.Error(), "EmailWithValidValue")
	}

	return value
}

func EmailWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewEmail(value)

	return value, err
}
