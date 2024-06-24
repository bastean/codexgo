package valueobjs

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func EntityWithValidValue() models.ValueObject[string] {
	value, err := NewEntity(services.Create.Regex(`^[A-Za-z]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "EntityWithValidValue")
	}

	return value
}

func EntityWithInvalidLength() (string, error) {
	value := ""

	_, err := NewEntity(value)

	return value, err
}

func EntityWithInvalidAlpha() (string, error) {
	value := "<></>"

	_, err := NewEntity(value)

	return value, err
}
