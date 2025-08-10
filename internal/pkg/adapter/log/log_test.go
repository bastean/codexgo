package log_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/internal/pkg/adapter/log"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type LogTestSuite struct {
	suite.Default
}

func (s *LogTestSuite) TestSentinel() {
	s.Equal("speed", log.FontName)
	s.Equal(5, log.FontHeight)
}

func TestUnitLogSuite(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}
