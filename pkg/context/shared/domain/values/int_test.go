package values_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type IntTestSuite struct {
	suite.Default
}

func (s *IntTestSuite) TestIntPositiveWithInvalidValue() {
	value, err := values.Mother().IntPositiveInvalid()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Validate",
		What:  "Invalid positive number",
		Why: errors.Meta{
			"Number": value,
		},
	}}

	s.Equal(expected, actual)
}

func (s *IntTestSuite) TestIntNegativeWithInvalidValue() {
	value, err := values.Mother().IntNegativeInvalid()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Validate",
		What:  "Invalid negative number",
		Why: errors.Meta{
			"Number": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitIntSuite(t *testing.T) {
	suite.Run(t, new(IntTestSuite))
}
