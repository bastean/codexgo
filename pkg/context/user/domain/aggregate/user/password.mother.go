package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

func PlainPasswordWithValidValue() *PlainPassword {
	value, err := values.New[*PlainPassword](services.Create.Regex(`^[\W\w]{8,64}$`))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func PlainPasswordWithInvalidLength() (string, error) {
	value := "x"

	_, err := values.New[*PlainPassword](value)

	return value, err
}

func CipherPasswordWithValidValue() *CipherPassword {
	value, err := values.New[*CipherPassword](services.Create.Regex(`^.{8,64}$`))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func CipherPasswordWithInvalidValue() (string, error) {
	var value string

	_, err := values.New[*CipherPassword](value)

	return value, err
}
