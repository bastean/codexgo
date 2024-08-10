package valueobjs

import (
	"fmt"
	"strings"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

const (
	OrganizationMinCharactersLength = "1"
	OrganizationMaxCharactersLength = "20"
)

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
			What:  fmt.Sprintf("Organization must be between %s to %s characters and be alphanumeric only", OrganizationMinCharactersLength, OrganizationMaxCharactersLength),
			Why: errors.Meta{
				"Organization": value,
			},
		})
	}

	return valueObj, nil
}
