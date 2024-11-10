package components

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func ResponseWithValidValue() *Response {
	value, err := NewResponse(services.Create.Regex(`^[A-Za-z]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "ResponseWithValidValue")
	}

	return value
}

func ResponseWithInvalidLength() (string, error) {
	var value string

	_, err := NewResponse(value)

	return value, err
}

func ResponseWithInvalidAlpha() (string, error) {
	value := "<></>"

	_, err := NewResponse(value)

	return value, err
}
