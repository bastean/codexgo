package time_test

import (
	"testing"
	t "time"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type WrapTestSuite struct {
	suite.Default
}

func (s *WrapTestSuite) TestWrapped() {
	s.Equal(time.RFC3339Nano, t.RFC3339Nano)

	s.Equal(time.Millisecond, t.Millisecond)
	s.Equal(time.Second, t.Second)
	s.Equal(time.Minute, t.Minute)
	s.Equal(time.Hour, t.Hour)

	s.Equal(time.Day, time.Hour*24)
	s.Equal(time.Week, time.Day*7)
	s.Equal(time.Month, time.Week*4)
	s.Equal(time.Year, time.Month*12)

	s.Equal(time.Duration(0), t.Duration(0))
}

func TestUnitWrapSuite(t *testing.T) {
	suite.Run(t, new(WrapTestSuite))
}
