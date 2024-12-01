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

func (s *CommandTestSuite) TestWithInvalidLength() {
	value, err := components.CommandWithInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "Command must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	s.Equal(expected, actual)
}

func (s *CommandTestSuite) TestWithInvalidAlpha() {
	value, err := components.CommandWithInvalidAlpha()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCommand",
		What:  "Command must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Command": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitCommandSuite(t *testing.T) {
	suite.Run(t, new(CommandTestSuite))
}
