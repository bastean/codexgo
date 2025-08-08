package errors_test

import (
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type AssertionTestSuite struct {
	suite.Default
}

func (s *AssertionTestSuite) TestErrTypeAssertion() {
	what := errors.Mother().Word()

	err := errors.Assertion(what)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		ID:    actual.ID,
		When:  actual.When,
		Where: "errors_test/*AssertionTestSuite/TestErrTypeAssertion",
		What:  fmt.Sprintf("Failure in %s type assertion", what),
	}}

	s.Equal(expected, actual)
}

func TestUnitAssertionSuite(t *testing.T) {
	suite.Run(t, new(AssertionTestSuite))
}
