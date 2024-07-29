package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func PasswordWithValidValue() *Password {
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
