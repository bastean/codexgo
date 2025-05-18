package errors_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type BubbleTestSuite struct {
	suite.Suite
}

func (s *BubbleTestSuite) TestWithValidValue() {
	bubble := errors.Mother().BubbleValid()

	why, err := json.Marshal(bubble.Why)

	s.NoError(err)

	expected := fmt.Sprintf("%s (%s): %s: %s: [%s]",
		bubble.When.Format(),
		bubble.Where,
		bubble.What,
		why,
		bubble.Who,
	)

	err = errors.New[errors.Default](bubble)

	var actual *errors.Default

	s.ErrorAs(err, &actual)

	s.Equal(expected, actual.Error())
}

func (s *BubbleTestSuite) TestWithoutWhere() {
	expected := "(New): Cannot create a error Bubble if \"Where\" is not defined"
	s.PanicsWithValue(expected, func() { errors.Mother().BubbleInvalidWithoutWhere() })
}

func (s *BubbleTestSuite) TestWithoutWhat() {
	expected := "(New): Cannot create a error Bubble if \"What\" is not defined"
	s.PanicsWithValue(expected, func() { errors.Mother().BubbleInvalidWithoutWhat() })
}

func TestUnitBubbleSuite(t *testing.T) {
	suite.Run(t, new(BubbleTestSuite))
}
