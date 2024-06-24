package valueobjs

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func ActionWithValidValue() models.ValueObject[string] {
	value, err := NewAction(services.Create.Regex(`^[A-Za-z\s]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "ActionWithValidValue")
	}

	return value
}

func ActionWithInvalidLength() (string, error) {
	value := ""

	_, err := NewAction(value)

	return value, err
}
