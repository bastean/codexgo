package errors_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type UnwrapTestSuite struct {
	suite.Default
}

func (s *UnwrapTestSuite) TestSentinel() {
	s.EqualValues(new(struct {
		Internal     []*errors.Internal
		Failure      []*errors.Failure
		InvalidValue []*errors.InvalidValue
		AlreadyExist []*errors.AlreadyExist
		NotExist     []*errors.NotExist
		Unknown      []error
		Amount       int
	}), new(errors.Bubbles))
}

func (s *UnwrapTestSuite) TestExtractBubbles() {
	var (
		internal     = errors.Mother().InternalValid()
		failure      = errors.Mother().FailureValid()
		invalidValue = errors.Mother().InvalidValueValid()
		alreadyExist = errors.Mother().AlreadyExistValid()
		notExist     = errors.Mother().NotExistValid()
		unknown      = errors.Mother().Error()
	)

	joined := errors.Join(internal, failure, invalidValue, alreadyExist, notExist, unknown)

	wrapped := errors.BubbleUp(joined)

	actual := errors.ExtractBubbles(wrapped)

	expected := []error{internal, failure, invalidValue, alreadyExist, notExist, unknown}

	s.Equal(expected, actual)
}

func (s *UnwrapTestSuite) TestFilterBubbles() {
	var (
		internal     = errors.Mother().InternalValid()
		failure      = errors.Mother().FailureValid()
		invalidValue = errors.Mother().InvalidValueValid()
		alreadyExist = errors.Mother().AlreadyExistValid()
		notExist     = errors.Mother().NotExistValid()
		unknown      = errors.Mother().Error()
	)

	actual := new(errors.Bubbles)

	errors.FilterBubbles([]error{internal, failure, invalidValue, alreadyExist, notExist, unknown}, actual)

	expected := &errors.Bubbles{
		Internal:     []*errors.Internal{internal},
		Failure:      []*errors.Failure{failure},
		InvalidValue: []*errors.InvalidValue{invalidValue},
		AlreadyExist: []*errors.AlreadyExist{alreadyExist},
		NotExist:     []*errors.NotExist{notExist},
		Unknown:      []error{unknown},
		Amount:       6,
	}

	s.Equal(expected, actual)
}

func (s *UnwrapTestSuite) TestFilterBubblesErrBubblesNotDefined() {
	expected := "(errors/FilterBubbles): Cannot filter if \"Bubbles\" are not defined"
	s.PanicsWithValue(expected, func() { errors.FilterBubbles(make([]error, 0), nil) })
}

func (s *UnwrapTestSuite) TestUnwrap() {
	var (
		internal     = errors.Mother().InternalValid()
		failure      = errors.Mother().FailureValid()
		invalidValue = errors.Mother().InvalidValueValid()
		alreadyExist = errors.Mother().AlreadyExistValid()
		notExist     = errors.Mother().NotExistValid()
		unknown      = errors.Mother().Error()
	)

	joined := errors.Join(internal, failure, invalidValue, alreadyExist, notExist, unknown)

	wrapped := errors.BubbleUp(joined)

	actual := new(errors.Bubbles)

	errors.Unwrap(wrapped, actual)

	expected := &errors.Bubbles{
		Internal:     []*errors.Internal{internal},
		Failure:      []*errors.Failure{failure},
		InvalidValue: []*errors.InvalidValue{invalidValue},
		AlreadyExist: []*errors.AlreadyExist{alreadyExist},
		NotExist:     []*errors.NotExist{notExist},
		Unknown:      []error{unknown},
		Amount:       6,
	}

	s.Equal(expected, actual)
}

func TestUnitUnwrapSuite(t *testing.T) {
	suite.Run(t, new(UnwrapTestSuite))
}
