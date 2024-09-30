package components

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

func CommandWithValidValue() *Command {
	value, err := NewCommand(services.Create.Regex(`^[A-Za-z]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "CommandWithValidValue")
	}

	return value
}

func CommandWithInvalidLength() (string, error) {
	var value string

	_, err := NewCommand(value)

	return value, err
}

func CommandWithInvalidAlpha() (string, error) {
	value := "<></>"

	_, err := NewCommand(value)

	return value, err
}
