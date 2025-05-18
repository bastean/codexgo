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

	s.Equal(caller.DefaultPkg, "UNKNOWN")
	s.Equal(caller.DefaultReceiver, "UNKNOWN")
	s.Equal(caller.DefaultMethod, "UNKNOWN")
}

func (s *CallerTestSuite) TestCallerFromCurrent() {
	pkg, receiver, method := caller.Received(caller.FromCurrent)

	s.Equal(pkg, "caller_test")

	s.Equal(receiver, "*CallerTestSuite")

	s.Equal(method, "TestCallerFromCurrent")
}

func (s *CallerTestSuite) TestCallerWithSkipCurrent() {
	var pkg, receiver, method string

	func() {
		pkg, receiver, method = caller.Received(caller.SkipCurrent)
	}()

	s.Equal(pkg, "caller_test")

	s.Equal(receiver, "*CallerTestSuite")

	s.Equal(method, "TestCallerWithSkipCurrent")
}

func (s *CallerTestSuite) TestCallerWithAnonymous() {
	var pkg, receiver, method string

	func() {
		pkg, receiver, method = caller.Received(caller.FromCurrent)
	}()

	s.Equal(pkg, "UNKNOWN")

	s.Equal(receiver, "UNKNOWN")

	s.Equal(method, "UNKNOWN")
}

func TestUnitCallerSuite(t *testing.T) {
	suite.Run(t, new(CallerTestSuite))
}
