package errors_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/embed"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type BubbleTestSuite struct {
	suite.Default
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
	err := errors.Mother().BubbleValidWithoutWhere()

	var actual *errors.Default

	s.ErrorAs(err, &actual)

	expected := &errors.Default{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "UNKNOWN",
		What:  actual.What,
	}}

	s.Equal(expected, actual)
}

func (s *BubbleTestSuite) TestWithoutWhat() {
	expected := "(errors/New): Cannot create a error Bubble if \"What\" is not defined"
	s.PanicsWithValue(expected, func() { errors.Mother().BubbleInvalidWithoutWhat() })
}

func (s *BubbleTestSuite) TestWithInvalidWhy() {
	defer func() {
		if actual := recover(); s.NotNil(actual) {
			expected := fmt.Sprintf("(errors/*Bubble/Error): Cannot format \"Why\" from error Bubble [%s]", embed.Extract(actual.(string)))
			s.Equal(expected, actual)
		}
	}()

	errors.Mother().BubbleInvalidWhy()
}

func TestUnitBubbleSuite(t *testing.T) {
	suite.Run(t, new(BubbleTestSuite))
}
