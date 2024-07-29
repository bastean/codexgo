package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
	"github.com/stretchr/testify/suite"
)

type OrganizationValueObjectTestSuite struct {
	suite.Suite
}

func (suite *OrganizationValueObjectTestSuite) SetupTest() {}

func (suite *OrganizationValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.OrganizationWithInvalidLength()

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

func (suite *OrganizationValueObjectTestSuite) TestWithInvalidAlphanumeric() {
	value, err := valueobjs.OrganizationWithInvalidAlphanumeric()

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

func TestUnitOrganizationValueObjectSuite(t *testing.T) {
	suite.Run(t, new(OrganizationValueObjectTestSuite))
}
