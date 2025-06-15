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

func (s *WrapTestSuite) TestSentinel() {
	s.Equal(t.RFC3339Nano, time.RFC3339Nano)

	s.Equal(t.Millisecond, time.Millisecond)
	s.Equal(t.Second, time.Second)
	s.Equal(t.Minute, time.Minute)
	s.Equal(t.Hour, time.Hour)

	s.Equal(time.Hour*24, time.Day)
	s.Equal(time.Day*7, time.Week)
	s.Equal(time.Week*4, time.Month)
	s.Equal(time.Month*12, time.Year)

	s.Equal(t.Duration(0), time.Duration(0))
}

func TestUnitWrapSuite(t *testing.T) {
	suite.Run(t, new(WrapTestSuite))
}
