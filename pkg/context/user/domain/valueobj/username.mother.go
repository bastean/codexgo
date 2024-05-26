package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func RandomUsername() (models.ValueObject[string], error) {
	return NewUsername(services.Create.Regex(`[a-z0-9]{2,20}`))
}

func WithInvalidUsernameLength() (string, error) {
	value := "x"

	_, err := NewUsername(value)

	return value, err
}

func WithInvalidUsernameAlphanumeric() (string, error) {
	value := "<></>"

	_, err := NewUsername(value)

	return value, err
}
