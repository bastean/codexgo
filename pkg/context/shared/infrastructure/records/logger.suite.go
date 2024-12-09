package records

import (
	"bytes"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/loggers"
)

type LoggerSuite struct {
	suite.Suite
	SUT    loggers.Logger
	Buffer *bytes.Buffer
}

func (s *LoggerSuite) TestDebug() {
	message := loggers.RandomMessage()

	s.SUT.Debug(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestError() {
	message := loggers.RandomMessage()

	s.SUT.Error(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestInfo() {
	message := loggers.RandomMessage()

	s.SUT.Info(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestSuccess() {
	message := loggers.RandomMessage()

	s.SUT.Success(message)

	s.Contains(s.Buffer.String(), message)
}
