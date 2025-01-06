package aggregates_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type TimeTestSuite struct {
	suite.Suite
}

func (s *TimeTestSuite) TestWithInvalidValue() {
	value, err := aggregates.TimeWithInvalidValue()

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewTime",
		What:  "Invalid Time format",
		Why: errors.Meta{
			"Time": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitTimeSuite(t *testing.T) {
	suite.Run(t, new(TimeTestSuite))
}
