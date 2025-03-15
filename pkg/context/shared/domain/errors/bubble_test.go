package errors_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services"
)

type BubbleTestSuite struct {
	suite.Suite
}

func (s *BubbleTestSuite) TestWithValidValue() {
	errThirdParty := services.Create.Error()

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

func (s *BubbleTestSuite) TestWithInvalidValue() {
	s.Panics(func() { _ = errors.New[errors.Default](new(errors.Bubble)) })
}

func TestUnitBubbleSuite(t *testing.T) {
	suite.Run(t, new(BubbleTestSuite))
}
