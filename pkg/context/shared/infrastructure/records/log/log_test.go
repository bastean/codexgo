package log_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/records/log"
)

type LogTestSuite struct {
	records.LoggerSuite
}

func (s *LogTestSuite) SetupTest() {
	logger := log.New()

	s.LoggerSuite.Buffer = new(bytes.Buffer)

	logger.SetOutput(s.LoggerSuite.Buffer)

	s.LoggerSuite.SUT = logger
}

func TestIntegrationLogSuite(t *testing.T) {
	suite.Run(t, new(LogTestSuite))
}
