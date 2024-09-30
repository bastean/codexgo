package components

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func ActionWithValidValue() *Action {
	value, err := NewAction(services.Create.Regex(`^[A-Za-z\s]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "ActionWithValidValue")
	}

	return value
}

func ActionWithInvalidLength() (string, error) {
	var value string

	_, err := NewAction(value)

	return value, err
}
