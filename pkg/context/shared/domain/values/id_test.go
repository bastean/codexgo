package values_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type IDTestSuite struct {
	suite.Default
}

func (s *IDTestSuite) TestErrInvalidFormat() {
	value, err := values.Mother().IDInvalid()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "values/*ID/Validate",
		What:  "Invalid UUID4 format",
		Why: errors.Meta{
			"ID": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitIDSuite(t *testing.T) {
	suite.Run(t, new(IDTestSuite))
}
