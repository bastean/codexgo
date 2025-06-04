package array_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/array"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type SliceTestSuite struct {
	suite.Default
}

func (s *SliceTestSuite) TestSlice() {
	values, index := array.Mother().SliceValid()

	actual, exists := array.Slice(values, index)

	s.True(exists)

	expected := values[index]

	s.Equal(expected, actual)
}

func (s *SliceTestSuite) TestSliceErrOutOfRangeByIndexMajor() {
	values, _ := array.Mother().SliceValid()

	actual, exists := array.Slice(values, len(values)+1)

	s.False(exists)

	expected := ""

	s.Equal(expected, actual)
}

func (s *SliceTestSuite) TestSliceErrOutOfRangeByIndexEqual() {
	values, _ := array.Mother().SliceValid()

	actual, exists := array.Slice(values, len(values))

	s.False(exists)

	expected := ""

	s.Equal(expected, actual)
}

func (s *SliceTestSuite) TestSliceErrOutOfRangeByIndexMinor() {
	values, _ := array.Mother().SliceValid()

	actual, exists := array.Slice(values, -len(values))

	s.False(exists)

	expected := ""

	s.Equal(expected, actual)
}

func TestUnitSliceSuite(t *testing.T) {
	suite.Run(t, new(SliceTestSuite))
}
