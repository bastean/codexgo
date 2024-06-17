package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func EmailWithValidValue() models.ValueObject[string] {
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
