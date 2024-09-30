package components

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func ServiceWithValidValue() *Service {
	value, err := NewService(services.Create.Regex(`^[A-Za-z0-9]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "ServiceWithValidValue")
	}

	return value
}

func ServiceWithInvalidLength() (string, error) {
	var value string

	_, err := NewService(value)

	return value, err
}

func ServiceWithInvalidAlphanumeric() (string, error) {
	value := "<></>"

	_, err := NewService(value)

	return value, err
}
