package aggregates

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func TimeWithValidValue() *Time {
	value, err := NewTime(services.TimeNow())

	if err != nil {
		errors.Panic(err.Error(), "TimeWithValidValue")
	}

	return value
}

func TimeWithInvalidValue() (string, error) {
	var value string

	_, err := NewTime(value)

	return value, err
}
