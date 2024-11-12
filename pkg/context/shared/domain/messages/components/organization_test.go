package components_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
	"github.com/stretchr/testify/suite"
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

	s.EqualError(expected, actual.Error())
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

	s.EqualError(expected, actual.Error())
}

func TestUnitOrganizationSuite(t *testing.T) {
	suite.Run(t, new(OrganizationTestSuite))
}
