package valueobj

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func UsernameWithValidValue() models.ValueObject[string] {
	value, err := NewUsername(services.Create.Regex(`^[A-Za-z0-9]{2,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "UsernameWithValidValue")
	}

	return value
}

func UsernameWithInvalidLength() (string, error) {
	value := "x"

	_, err := NewUsername(value)

	return value, err
}

func UsernameWithInvalidAlphanumeric() (string, error) {
	value := "<></>"

	_, err := NewUsername(value)

	return value, err
}
