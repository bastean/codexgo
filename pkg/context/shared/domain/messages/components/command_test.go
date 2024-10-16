package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type CommandTestSuite struct {
	suite.Suite
}

func (suite *CommandTestSuite) TestWithInvalidLength() {
	value, err := components.CommandWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "Command must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CommandTestSuite) TestWithInvalidAlpha() {
	value, err := components.CommandWithInvalidAlpha()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "Command must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}
