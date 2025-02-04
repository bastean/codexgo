package errors_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
)

type BubbleTestSuite struct {
	suite.Suite
}

func (s *BubbleTestSuite) TestWithValidValue() {
	bubble := &errors.Bubble{
		When:  time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		Where: "TestWithValidValue",
		What:  "Test Case",
		Why: errors.Meta{
			"Case": "Happy path",
		},
		Who: fmt.Errorf("third-party error"),
	}

	expected := "2009-11-10T23:00:00Z (TestWithValidValue): Test Case: {\"Case\":\"Happy path\"}: [third-party error]"

	err := errors.New[errors.Default](bubble)

	var actual *errors.Default

	s.ErrorAs(err, &actual)

	s.Equal(expected, actual.Error())
}

func (s *BubbleTestSuite) TestWithInvalidValue() {
	s.Panics(func() { errors.New[errors.Default](new(errors.Bubble)) })
}

func TestUnitBubbleSuite(t *testing.T) {
	suite.Run(t, new(BubbleTestSuite))
}
