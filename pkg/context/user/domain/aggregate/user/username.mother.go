package user

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

func UsernameWithValidValue() *Username {
	value, err := values.New[*Username](services.Create.Regex(`^[A-Za-z0-9]{2,20}$`))

	if err != nil {
		errors.Panic(err)
	}

	return value
}

func UsernameWithInvalidLength() (string, error) {
	value := "x"

	_, err := values.New[*Username](value)

	return value, err
}

func UsernameWithInvalidAlphanumeric() (string, error) {
	value := "<></>"

	_, err := values.New[*Username](value)

	return value, err
}
