package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type OrganizationTestSuite struct {
	suite.Suite
}

func (s *OrganizationTestSuite) TestWithInvalidLength() {
	value, err := components.OrganizationWithInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewOrganization",
		What:  "Organization must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Organization": value,
		},
	}}

	s.Equal(expected, actual)
}

func (s *OrganizationTestSuite) TestWithInvalidAlphanumeric() {
	value, err := components.OrganizationWithInvalidAlphanumeric()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewOrganization",
		What:  "Organization must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Organization": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitOrganizationSuite(t *testing.T) {
	suite.Run(t, new(OrganizationTestSuite))
}
