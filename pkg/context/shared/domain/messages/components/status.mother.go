package components

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func StatusWithValidValue() *Status {
	value, err := NewStatus(services.Create.RandomString([]string{"queued", "succeeded", "failed", "done"}))

	if err != nil {
		errors.Panic(err.Error(), "StatusWithValidValue")
	}

	return value
}

func StatusWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewStatus(value)

	return value, err
}
