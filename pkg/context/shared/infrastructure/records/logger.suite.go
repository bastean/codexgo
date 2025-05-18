package records

import (
	"bytes"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type LoggerSuite struct {
	suite.Suite
	SUT    roles.Logger
	Buffer *bytes.Buffer
}

func (s *LoggerSuite) TestDebug() {
	message := Mother().Message()

	s.SUT.Debug(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestError() {
	message := Mother().Message()

	s.SUT.Error(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestInfo() {
	message := Mother().Message()

	s.SUT.Info(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestSuccess() {
	message := Mother().Message()

	s.SUT.Success(message)

	s.Contains(s.Buffer.String(), message)
}
