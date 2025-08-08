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

func (s *IntTestSuite) TestIntPositiveErrInvalidNumber() {
	value, err := values.Mother().IntPositiveInvalid()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: "values/*IntPositive/Validate",
		What:  "Invalid positive number",
		Why: errors.Meta{
			"Number": value,
		},
	}}

	s.Equal(expected, actual)
}

func (s *IntTestSuite) TestIntNegativeErrInvalidNumber() {
	value, err := values.Mother().IntNegativeInvalid()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: "values/*IntNegative/Validate",
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
