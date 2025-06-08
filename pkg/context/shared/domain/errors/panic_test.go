package errors_test

import (
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type PanicTestSuite struct {
	suite.Default
}

func (s *PanicTestSuite) TestPanic() {
	err := errors.Mother().Error()

	expected := fmt.Sprintf("(errors/*m/PanicValidWithError): %s", err)

	s.PanicsWithValue(expected, func() { errors.Mother().PanicValidWithError(err) })
}

func (s *PanicTestSuite) TestPanicWithAnonymous() {
	err := errors.Mother().Error()

	expected := fmt.Sprintf("(UNKNOWN): %s", err)

	s.PanicsWithValue(expected, func() { errors.Panic(err) })
}

func TestUnitPanicSuite(t *testing.T) {
	suite.Run(t, new(PanicTestSuite))
}
