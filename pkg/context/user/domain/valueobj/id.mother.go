package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func RandomId() (models.ValueObject[string], error) {
	return NewId(services.Create.UUID())
}

func InvalidId() (string, error) {
	value := "x"

	_, err := NewId(value)

	return value, err
}
