package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
	"github.com/stretchr/testify/suite"
)

type EventValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EventValueObjectTestSuite) SetupTest() {}

func (suite *EventValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.EventWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEvent",
		What:  "Event must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Event": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *EventValueObjectTestSuite) TestWithInvalidAlpha() {
	value, err := valueobjs.EventWithInvalidAlpha()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEvent",
		What:  "Event must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Event": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEventValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EventValueObjectTestSuite))
}
