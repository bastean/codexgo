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

func (s *ServiceTestSuite) TestWithInvalidLength() {
	value, err := components.ServiceWithInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewService",
		What:  "Service must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Service": value,
		},
	}}

	s.Equal(expected, actual)
}

func (s *ServiceTestSuite) TestWithInvalidAlphanumeric() {
	value, err := components.ServiceWithInvalidAlphanumeric()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewService",
		What:  "Service must be between 1 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Service": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitServiceSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
