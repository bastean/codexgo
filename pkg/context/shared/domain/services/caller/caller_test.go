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

func (s *CallerTestSuite) TestCallerWithFromCurrent() {
	where, pkg, receiver, method := caller.Received(caller.FromCurrent)

	s.Equal("caller_test/*CallerTestSuite/TestCallerWithFromCurrent", where)

	s.Equal("caller_test", pkg)

	s.Equal("*CallerTestSuite", receiver)

	s.Equal("TestCallerWithFromCurrent", method)
}

func (s *CallerTestSuite) TestCallerWithSkipCurrent() {
	var where, pkg, receiver, method string

	func() {
		where, pkg, receiver, method = caller.Received(caller.SkipCurrent)
	}()

	s.Equal("caller_test/*CallerTestSuite/TestCallerWithSkipCurrent", where)

	s.Equal("caller_test", pkg)

	s.Equal("*CallerTestSuite", receiver)

	s.Equal("TestCallerWithSkipCurrent", method)
}

func (s *CallerTestSuite) TestCallerWithAnonymousFunction() {
	var where, pkg, receiver, method string

	func() {
		where, pkg, receiver, method = caller.Received(caller.FromCurrent)
	}()

	s.Equal("UNKNOWN", where)

	s.Equal("UNKNOWN", pkg)

	s.Equal("UNKNOWN", receiver)

	s.Equal("UNKNOWN", method)
}

func TestUnitCallerSuite(t *testing.T) {
	suite.Run(t, new(CallerTestSuite))
}
