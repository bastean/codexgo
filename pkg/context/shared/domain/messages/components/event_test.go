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

func (suite *EventTestSuite) TestWithInvalidLength() {
	value, err := components.EventWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEvent",
		What:  "Event must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Event": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *EventTestSuite) TestWithInvalidAlpha() {
	value, err := components.EventWithInvalidAlpha()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEvent",
		What:  "Event must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Event": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEventSuite(t *testing.T) {
	suite.Run(t, new(EventTestSuite))
}
