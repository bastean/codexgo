package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func RandomEmail() (models.ValueObject[string], error) {
	return NewEmail(services.Create.Email())
}

func InvalidEmail() (string, error) {
	value := "x"

	_, err := NewEmail(value)

	return value, err
}
