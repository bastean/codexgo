package valueobjs

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func TypeWithValidValue() models.ValueObject[string] {
	value, err := NewType(services.Create.RandomString([]string{"event", "command"}))

	if err != nil {
		errors.Panic(err.Error(), "TypeWithValidValue")
	}

	return value
}

func TypeWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewType(value)

	return value, err
}
