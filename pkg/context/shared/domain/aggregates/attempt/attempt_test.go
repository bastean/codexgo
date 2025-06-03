package attempt_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/attempt"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
)

type AttemptTestSuite struct {
	suite.Frozen
}

func (s *AttemptTestSuite) TestIncrease() {
	aggregate := attempt.Mother().AttemptValid()

	for i := range aggregate.Limit.Value() {
		s.SetTimeAfter(time.Duration(aggregate.Every.Value() * i))
		s.NoError(aggregate.Increase())
	}

	expected := strings.Split(time.Now().Add(time.Duration(aggregate.Next.Value())).Format(), ".")[0]

	actual := strings.Split(aggregate.Until.Value(), ".")[0]

	s.Equal(expected, actual)

	s.Equal(aggregate.Counter.Value(), 0)
}

func (s *AttemptTestSuite) TestIncreaseErrNoMore() {
	aggregate := attempt.Mother().AttemptValid()

	for i := range aggregate.Limit.Value() {
		s.SetTimeAfter(time.Duration(aggregate.Every.Value() * i))
		s.NoError(aggregate.Increase())
	}

	err := aggregate.Increase()

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Increase",
		What:  fmt.Sprintf("No more attempts, please try again in %q", time.Now().Sub(time.Parse(aggregate.Until.Value())).Round(time.Second)),
	}}

	s.Equal(expected, actual)
}

func (s *AttemptTestSuite) TestIncreaseErrTryAgain() {
	aggregate := attempt.Mother().AttemptValid()

	current := time.Now()

	aggregate.Counter.SetUpdatedAt(attempt.Mother().TimeSetAfter(current, time.Hour, time.Day))

	err := aggregate.Increase()

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Increase",
		What:  fmt.Sprintf("Please try again in %q", current.Sub(time.Parse(aggregate.Counter.UpdatedAt()).Add(time.Duration(aggregate.Every.Value()))).Round(time.Second)),
	}}

	s.Equal(expected, actual)
}

func TestUnitAttemptSuite(t *testing.T) {
	suite.Run(t, new(AttemptTestSuite))
}
