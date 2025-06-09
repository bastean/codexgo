package values_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type UsernameTestSuite struct {
	suite.Default
}

func (s *UsernameTestSuite) TestErrInvalidLength() {
	value, err := values.Mother().UsernameInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "values/*Username/Validate",
		What:  "Username must be between 2 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Username": value,
		},
	}}

	s.Equal(expected, actual)
}

func (s *UsernameTestSuite) TestErrInvalidAlphanumeric() {
	value, err := values.Mother().UsernameInvalidAlphanumeric()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "values/*Username/Validate",
		What:  "Username must be between 2 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Username": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitUsernameSuite(t *testing.T) {
	suite.Run(t, new(UsernameTestSuite))
}
