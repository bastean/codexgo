package log_test

import (
	"bytes"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
)

type LogTestSuite struct {
	records.LoggerSuite
}

func (s *LogTestSuite) SetupSuite() {
	logger := log.New()

	s.LoggerSuite.Actual = new(bytes.Buffer)

	logger.SetOutput(s.LoggerSuite.Actual)

	s.LoggerSuite.SUT = logger
}

func TestIntegrationLogSuite(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}
