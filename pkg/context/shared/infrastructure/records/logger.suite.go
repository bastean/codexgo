package records

import (
	"bytes"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type LoggerSuite struct {
	suite.Default
	SUT    roles.Logger
	Actual *bytes.Buffer
}

func (s *LoggerSuite) TestDebug() {
	format, values, expected := Mother().Message()

	s.SUT.Debug(format, values...)

	s.Contains(s.Actual.String(), expected)
}

func (s *LoggerSuite) TestError() {
	format, values, expected := Mother().Message()

	s.SUT.Error(format, values...)

	s.Contains(s.Actual.String(), expected)
}

func (s *LoggerSuite) TestInfo() {
	format, values, expected := Mother().Message()

	s.SUT.Info(format, values...)

	s.Contains(s.Actual.String(), expected)
}

func (s *LoggerSuite) TestSuccess() {
	format, values, expected := Mother().Message()

	s.SUT.Success(format, values...)

	s.Contains(s.Actual.String(), expected)
}
