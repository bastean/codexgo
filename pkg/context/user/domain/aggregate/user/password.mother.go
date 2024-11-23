package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func PlainPasswordWithValidValue() *PlainPassword {
	value, err := NewPlainPassword(services.Create.Regex(`^[\W\w]{8,64}$`))

	if err != nil {
		errors.Panic(err.Error(), "PlainPasswordWithValidValue")
	}

	return value
}

func PlainPasswordWithInvalidLength() (string, error) {
	value := "x"

	_, err := NewPlainPassword(value)

	return value, err
}

func CipherPasswordWithValidValue() *CipherPassword {
	value, err := NewCipherPassword(services.Create.Regex(`^.{8,64}$`))

	if err != nil {
		errors.Panic(err.Error(), "CipherPasswordWithValidValue")
	}

	return value
}

func CipherPasswordWithInvalidValue() (string, error) {
	var value string

	_, err := NewCipherPassword(value)

	return value, err
}
