package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type PanicTestSuite struct {
	suite.Suite
}

func (s *PanicTestSuite) TestPanic() {
	err := errors.Mother().Error()

	expected := fmt.Sprintf("(PanicValidWithError): %s", err)

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
