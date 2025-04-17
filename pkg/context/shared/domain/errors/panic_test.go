package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type PanicTestSuite struct {
	suite.Suite
}

func (s *PanicTestSuite) TestPanic() {
	err := services.Create.Error()

	expected := fmt.Sprintf("(PanicWithRandomValue): %s", err)

	s.PanicsWithValue(expected, func() { errors.PanicWithRandomValue(err) })
}

func (s *PanicTestSuite) TestPanicWithUnknown() {
	err := services.Create.Error()

	expected := fmt.Sprintf("(Unknown): %s", err)

	s.PanicsWithValue(expected, func() { errors.Panic(err) })
}

func TestUnitPanicSuite(t *testing.T) {
	suite.Run(t, new(PanicTestSuite))
}
