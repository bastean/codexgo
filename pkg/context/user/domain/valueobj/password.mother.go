package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func PasswordWithValidValue() (models.ValueObject[string], error) {
	return NewPassword(services.Create.Regex(`^[\W\w]{8,64}$`))
}

func PasswordWithInvalidLength() (string, error) {
	value := "x"

	_, err := NewPassword(value)

	return value, err
}
