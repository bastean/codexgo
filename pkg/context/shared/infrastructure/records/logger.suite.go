package records

import (
	"bytes"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type LoggerSuite struct {
	suite.Suite
	SUT    roles.Logger
	Buffer *bytes.Buffer
}

func (s *LoggerSuite) TestDebug() {
	message := services.Create.Message()

	s.SUT.Debug(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestError() {
	message := services.Create.Message()

	s.SUT.Error(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestInfo() {
	message := services.Create.Message()

	s.SUT.Info(message)

	s.Contains(s.Buffer.String(), message)
}

func (s *LoggerSuite) TestSuccess() {
	message := services.Create.Message()

	s.SUT.Success(message)

	s.Contains(s.Buffer.String(), message)
}
