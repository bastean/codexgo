package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type EventTestSuite struct {
	suite.Suite
}

func (s *EventTestSuite) TestWithInvalidLength() {
	value, err := components.EventWithInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEvent",
		What:  "Event must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Event": value,
		},
	}}

	s.Equal(expected, actual)
}

func (s *EventTestSuite) TestWithInvalidAlpha() {
	value, err := components.EventWithInvalidAlpha()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEvent",
		What:  "Event must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Event": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitEventSuite(t *testing.T) {
	suite.Run(t, new(EventTestSuite))
}
