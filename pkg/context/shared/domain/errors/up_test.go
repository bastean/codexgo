package errors_test

import (
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type BubbleUpTestSuite struct {
	suite.Default
}

func (s *BubbleUpTestSuite) TestBubbleUp() {
	err, value := errors.Mother().BubbleUpValid()

	actual := errors.BubbleUp(err)

	expected := fmt.Sprintf("(errors_test/*BubbleUpTestSuite/TestBubbleUp): [(errors/*m/BubbleUpValid): [%s]]", value)

	s.Equal(expected, actual.Error())
}

func (s *BubbleUpTestSuite) TestBubbleUpWithAnonymous() {
	var actual error

	err := errors.Mother().Error()

	func() {
		actual = errors.BubbleUp(err)
	}()

	expected := fmt.Sprintf("(UNKNOWN): [%s]", err)

	s.Equal(expected, actual.Error())
}

func TestUnitBubbleUpSuite(t *testing.T) {
	suite.Run(t, new(BubbleUpTestSuite))
}
