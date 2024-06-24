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

func (organization *Organization) Value() string {
	return organization.Organization
}

func (organization *Organization) IsValid() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return validate.Struct(organization)
}

func NewOrganization(organization string) (models.ValueObject[string], error) {
	organization = strings.TrimSpace(organization)

	organizationVO := &Organization{
		Organization: organization,
	}

	if organizationVO.IsValid() != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewOrganization",
			What:  "organization must be between " + OrganizationMinCharactersLength + " to " + OrganizationMaxCharactersLength + " characters and be alphanumeric only",
			Why: errors.Meta{
				"Organization": organization,
			},
		})
	}

	return organizationVO, nil
}
