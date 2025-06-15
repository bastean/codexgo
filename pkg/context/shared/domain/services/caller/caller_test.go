package caller_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/caller"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type CallerTestSuite struct {
	suite.Default
}

func (s *CallerTestSuite) TestSentinel() {
	s.Equal(0, caller.FromCurrent)
	s.Equal(1, caller.SkipCurrent)

	s.Equal("/", caller.Separator)

	s.Equal("UNKNOWN", caller.DefaultWhere)
	s.Equal("UNKNOWN", caller.DefaultPkg)
	s.Equal("UNKNOWN", caller.DefaultReceiver)
	s.Equal("UNKNOWN", caller.DefaultMethod)
}

func (s *CallerTestSuite) TestCallerFromCurrent() {
	where, pkg, receiver, method := caller.Received(caller.FromCurrent)

	s.Equal(where, "caller_test/*CallerTestSuite/TestCallerFromCurrent")

	s.Equal(pkg, "caller_test")

	s.Equal(receiver, "*CallerTestSuite")

	s.Equal(method, "TestCallerFromCurrent")
}

func (s *CallerTestSuite) TestCallerWithSkipCurrent() {
	var where, pkg, receiver, method string

	func() {
		where, pkg, receiver, method = caller.Received(caller.SkipCurrent)
	}()

	s.Equal(where, "caller_test/*CallerTestSuite/TestCallerWithSkipCurrent")

	s.Equal(pkg, "caller_test")

	s.Equal(receiver, "*CallerTestSuite")

	s.Equal(method, "TestCallerWithSkipCurrent")
}

func (s *CallerTestSuite) TestCallerWithAnonymous() {
	var where, pkg, receiver, method string

	func() {
		where, pkg, receiver, method = caller.Received(caller.FromCurrent)
	}()

	s.Equal(where, "UNKNOWN")

	s.Equal(pkg, "UNKNOWN")

	s.Equal(receiver, "UNKNOWN")

	s.Equal(method, "UNKNOWN")
}

func TestUnitCallerSuite(t *testing.T) {
	suite.Run(t, new(CallerTestSuite))
}
