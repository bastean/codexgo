package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
	"github.com/stretchr/testify/suite"
)

type OrganizationValueObjectTestSuite struct {
	suite.Suite
}

func (suite *OrganizationValueObjectTestSuite) SetupTest() {}

func (suite *OrganizationValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.OrganizationWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewOrganization",
		What:  "organization must be between " + "1" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Organization": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *OrganizationValueObjectTestSuite) TestWithInvalidAlphanumeric() {
	value, err := valueobjs.OrganizationWithInvalidAlphanumeric()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewOrganization",
		What:  "organization must be between " + "1" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Organization": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitOrganizationValueObjectSuite(t *testing.T) {
	suite.Run(t, new(OrganizationValueObjectTestSuite))
}
