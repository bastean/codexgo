package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func EmailWithValidValue() (models.ValueObject[string], error) {
	return NewEmail(services.Create.Email())
}

func EmailWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewEmail(value)

	return value, err
}
