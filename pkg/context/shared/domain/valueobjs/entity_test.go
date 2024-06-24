package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
	"github.com/stretchr/testify/suite"
)

type EntityValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EntityValueObjectTestSuite) SetupTest() {}

func (suite *EntityValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.EntityWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEntity",
		What:  "entity must be between " + "1" + " to " + "20" + " characters and be alpha only",
		Why: errors.Meta{
			"Entity": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *EntityValueObjectTestSuite) TestWithInvalidAlpha() {
	value, err := valueobjs.EntityWithInvalidAlpha()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEntity",
		What:  "entity must be between " + "1" + " to " + "20" + " characters and be alpha only",
		Why: errors.Meta{
			"Entity": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEntityValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EntityValueObjectTestSuite))
}
