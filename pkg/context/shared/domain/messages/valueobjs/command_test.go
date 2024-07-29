package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
	"github.com/stretchr/testify/suite"
)

type CommandValueObjectTestSuite struct {
	suite.Suite
}

func (suite *CommandValueObjectTestSuite) SetupTest() {}

func (suite *CommandValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.CommandWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "Command must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *CommandValueObjectTestSuite) TestWithInvalidAlpha() {
	value, err := valueobjs.CommandWithInvalidAlpha()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "Command must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitCommandValueObjectSuite(t *testing.T) {
	suite.Run(t, new(CommandValueObjectTestSuite))
}
