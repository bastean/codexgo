package components

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func EventWithValidValue() *Event {
	value, err := NewEvent(services.Create.Regex(`^[A-Za-z]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "EventWithValidValue")
	}

	return value
}

func EventWithInvalidLength() (string, error) {
	var value string

	_, err := NewEvent(value)

	return value, err
}

func EventWithInvalidAlpha() (string, error) {
	value := "<></>"

	_, err := NewEvent(value)

	return value, err
}
