package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
	"github.com/stretchr/testify/suite"
)

type CommandValueObjectTestSuite struct {
	suite.Suite
}

func (suite *CommandValueObjectTestSuite) SetupTest() {}

func (suite *CommandValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.CommandWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "command must be between " + "1" + " to " + "20" + " characters and be alpha only",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CommandValueObjectTestSuite) TestWithInvalidAlpha() {
	value, err := valueobjs.CommandWithInvalidAlpha()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "command must be between " + "1" + " to " + "20" + " characters and be alpha only",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCommandValueObjectSuite(t *testing.T) {
	suite.Run(t, new(CommandValueObjectTestSuite))
}
