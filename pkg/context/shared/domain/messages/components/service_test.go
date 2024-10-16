package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type ServiceTestSuite struct {
	suite.Suite
}

func (suite *ServiceTestSuite) TestWithInvalidLength() {
	value, err := components.ServiceWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewService",
		What:  "Service must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Service": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *ServiceTestSuite) TestWithInvalidAlphanumeric() {
	value, err := components.ServiceWithInvalidAlphanumeric()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewService",
		What:  "Service must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Service": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
