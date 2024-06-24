package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
	"github.com/stretchr/testify/suite"
)

type EventValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EventValueObjectTestSuite) SetupTest() {}

func (suite *EventValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.EventWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEvent",
		What:  "event must be between " + "1" + " to " + "20" + " characters and be alpha only",
		Why: errors.Meta{
			"Event": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *EventValueObjectTestSuite) TestWithInvalidAlpha() {
	value, err := valueobjs.EventWithInvalidAlpha()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEvent",
		What:  "event must be between " + "1" + " to " + "20" + " characters and be alpha only",
		Why: errors.Meta{
			"Event": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEventValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EventValueObjectTestSuite))
}
