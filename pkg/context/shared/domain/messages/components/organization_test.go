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

func (suite *OrganizationTestSuite) SetupTest() {}

func (suite *OrganizationTestSuite) TestWithInvalidLength() {
	value, err := components.OrganizationWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewOrganization",
		What:  "Organization must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Organization": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *OrganizationTestSuite) TestWithInvalidAlphanumeric() {
	value, err := components.OrganizationWithInvalidAlphanumeric()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewOrganization",
		What:  "Organization must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Organization": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitOrganizationSuite(t *testing.T) {
	suite.Run(t, new(OrganizationTestSuite))
}
