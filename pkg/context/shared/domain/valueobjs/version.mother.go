package valueobjs

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func VersionWithValidValue() models.ValueObject[string] {
	value, err := NewVersion(services.Create.Regex(`^[0-9]{1,2}$`))

	if err != nil {
		errors.Panic(err.Error(), "VersionWithValidValue")
	}

	return value
}

func VersionWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewVersion(value)

	return value, err
}
