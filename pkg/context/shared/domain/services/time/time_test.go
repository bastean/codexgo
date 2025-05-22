package time_test

import (
	"testing"
	t "time"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type TimeTestSuite struct {
	suite.Default
}

func (s *TimeTestSuite) TestParseWithValidValue() {
	actual := time.Now().Format()

	s.NotPanics(func() { _ = time.Parse(actual) })

	date, err := t.Parse(t.RFC3339Nano, actual)

	s.NoError(err)

	expected := date.UTC().Format(t.RFC3339Nano)

	s.Equal(expected, actual)
}

func (s *TimeTestSuite) TestParseWithInvalidValue() {
	expected := "Time format is not valid"
	s.PanicsWithValue(expected, func() { _ = time.Parse(t.Now().Format(t.Layout)) })
}

func TestUnitTimeSuite(t *testing.T) {
	suite.Run(t, new(TimeTestSuite))
}
