package valueobjs

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func EventWithValidValue() models.ValueObject[string] {
	value, err := NewEvent(services.Create.Regex(`^[A-Za-z]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "EventWithValidValue")
	}

	return value
}

func EventWithInvalidLength() (string, error) {
	value := ""

	_, err := NewEvent(value)

	return value, err
}

func EventWithInvalidAlpha() (string, error) {
	value := "<></>"

	_, err := NewEvent(value)

	return value, err
}
