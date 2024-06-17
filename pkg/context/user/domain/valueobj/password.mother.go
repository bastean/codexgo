package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func PasswordWithValidValue() models.ValueObject[string] {
	value, err := NewPassword(services.Create.Regex(`^[\W\w]{8,64}$`))

	if err != nil {
		errors.Panic(err.Error(), "PasswordWithValidValue")
	}

	return value
}

func PasswordWithInvalidLength() (string, error) {
	value := "x"

	_, err := NewPassword(value)

	return value, err
}
