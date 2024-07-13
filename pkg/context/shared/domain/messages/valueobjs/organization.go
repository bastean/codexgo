package valueobjs

import (
	"strings"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/services"
)

const OrganizationMinCharactersLength = "1"
const OrganizationMaxCharactersLength = "20"

type Organization struct {
	Value string `validate:"gte=1,lte=20,alphanum"`
}

func NewOrganization(value string) (*Organization, error) {
	value = strings.TrimSpace(value)

	valueObj := &Organization{
		Value: value,
	}

	if services.IsValueObjectInvalid(valueObj) {
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
