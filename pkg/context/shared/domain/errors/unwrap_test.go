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

	actual := errors.FilterBubbles([]error{internal, failure, invalidValue, alreadyExist, notExist, unknown})

	expected := &errors.Bubbles{
		Internal:     []*errors.Internal{internal},
		Failure:      []*errors.Failure{failure},
		InvalidValue: []*errors.InvalidValue{invalidValue},
		AlreadyExist: []*errors.AlreadyExist{alreadyExist},
		NotExist:     []*errors.NotExist{notExist},
		Unknown:      []error{unknown},
	}

	s.Equal(expected, actual)
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

	actual := errors.Unwrap(wrapped)

	expected := &errors.Bubbles{
		Internal:     []*errors.Internal{internal},
		Failure:      []*errors.Failure{failure},
		InvalidValue: []*errors.InvalidValue{invalidValue},
		AlreadyExist: []*errors.AlreadyExist{alreadyExist},
		NotExist:     []*errors.NotExist{notExist},
		Unknown:      []error{unknown},
	}

	s.Equal(expected, actual)
}

func TestUnitUnwrapSuite(t *testing.T) {
	suite.Run(t, new(UnwrapTestSuite))
}
