package errors_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type BubbleTestSuite struct {
	suite.Suite
}

func (s *BubbleTestSuite) TestWithValidValue() {
	errThirdParty := errors.Mother.Error()

	bubble := &errors.Bubble{
		When:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Where: "TestWithValidValue",
		What:  "Test Case",
		Why: errors.Meta{
			"Case": "Happy path",
		},
		Who: errThirdParty,
	}

	expected := "2009-11-10T23:00:00Z (TestWithValidValue): Test Case: {\"Case\":\"Happy path\"}: [" + errThirdParty.Error() + "]"

	err := errors.New[errors.Default](bubble)

	var actual *errors.Default

	s.ErrorAs(err, &actual)

	s.Equal(expected, actual.Error())
}

func (s *BubbleTestSuite) TestWithoutWhere() {
	expected := "(New): Cannot create a error Bubble if \"Where\" is not defined"
	s.PanicsWithValue(expected, func() { _ = errors.New[errors.Default](errors.Mother.BubbleInvalidWithoutWhere()) })
}

func (s *BubbleTestSuite) TestWithoutWhat() {
	expected := "(New): Cannot create a error Bubble if \"What\" is not defined"
	s.PanicsWithValue(expected, func() { _ = errors.New[errors.Default](errors.Mother.BubbleInvalidWithoutWhat()) })
}

func TestUnitBubbleSuite(t *testing.T) {
	suite.Run(t, new(BubbleTestSuite))
}
