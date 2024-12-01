package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type TypeTestSuite struct {
	suite.Suite
}

func (s *TypeTestSuite) TestWithInvalidValue() {
	value, err := components.TypeWithInvalidValue()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewType",
		What:  "Type must be only one of these values: Event, Command, Query, Response",
		Why: errors.Meta{
			"Type": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitTypeSuite(t *testing.T) {
	suite.Run(t, new(TypeTestSuite))
}
