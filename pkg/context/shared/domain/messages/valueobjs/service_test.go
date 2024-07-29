package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
	"github.com/stretchr/testify/suite"
)

type ServiceValueObjectTestSuite struct {
	suite.Suite
}

func (suite *ServiceValueObjectTestSuite) SetupTest() {}

func (suite *ServiceValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.ServiceWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewService",
		What:  "Service must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Service": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *ServiceValueObjectTestSuite) TestWithInvalidAlphanumeric() {
	value, err := valueobjs.ServiceWithInvalidAlphanumeric()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewService",
		What:  "Service must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Service": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitServiceValueObjectSuite(t *testing.T) {
	suite.Run(t, new(ServiceValueObjectTestSuite))
}
