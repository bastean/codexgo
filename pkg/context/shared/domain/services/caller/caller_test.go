package caller_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
)

type CallerTestSuite struct {
	suite.Suite
}

func (s *CallerTestSuite) SetupSuite() {
	s.Equal(caller.FromCurrent, 0)
	s.Equal(caller.SkipCurrent, 1)
}

func (s *CallerTestSuite) TestCaller() {
	pkg, receiver, method := caller.Received(caller.FromCurrent)

	s.Equal(pkg, "caller_test")

	s.Equal(receiver, "*CallerTestSuite")

	s.Equal(method, "TestCaller")
}

func (s *CallerTestSuite) TestCallerWithSkip() {
	var pkg, receiver, method string

	func() {
		pkg, receiver, method = caller.Received(caller.SkipCurrent)
	}()

	s.Equal(pkg, "caller_test")

	s.Equal(receiver, "*CallerTestSuite")

	s.Equal(method, "TestCallerWithSkip")
}

func TestUnitCallerSuite(t *testing.T) {
	suite.Run(t, new(CallerTestSuite))
}
