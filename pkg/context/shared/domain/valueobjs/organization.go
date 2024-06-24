package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/go-playground/validator/v10"
)

const OrganizationMinCharactersLength = "1"
const OrganizationMaxCharactersLength = "20"

type Organization struct {
	Organization string `validate:"gte=1,lte=20,alphanum"`
}

func (value *Organization) Value() string {
	return value.Organization
}

func (value *Organization) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(value)
}

func NewOrganization(value string) (models.ValueObject[string], error) {
	value = strings.TrimSpace(value)

	valueObj := &Organization{
		Organization: value,
	}

	if valueObj.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewOrganization",
			What:  "organization must be between " + OrganizationMinCharactersLength + " to " + OrganizationMaxCharactersLength + " characters and be alphanumeric only",
			Why: errors.Meta{
				"Organization": value,
			},
		})
	}

	return valueObj, nil
}
