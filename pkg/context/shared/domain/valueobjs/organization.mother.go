package valueobjs

import (
	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

func OrganizationWithValidValue() models.ValueObject[string] {
	value, err := NewOrganization(services.Create.Regex(`^[A-Za-z0-9]{1,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "OrganizationWithValidValue")
	}

	return value
}

func OrganizationWithInvalidLength() (string, error) {
	value := ""

	_, err := NewOrganization(value)

	return value, err
}

func OrganizationWithInvalidAlphanumeric() (string, error) {
	value := "<></>"

	_, err := NewOrganization(value)

	return value, err
}
