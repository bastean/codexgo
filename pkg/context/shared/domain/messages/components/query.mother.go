package components

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func QueryWithValidValue() *Query {
	value, err := NewQuery(services.Create.Regex(`^[A-Za-z]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "QueryWithValidValue")
	}

	return value
}

func QueryWithInvalidLength() (string, error) {
	var value string

	_, err := NewQuery(value)

	return value, err
}

func QueryWithInvalidAlpha() (string, error) {
	value := "<></>"

	_, err := NewQuery(value)

	return value, err
}
