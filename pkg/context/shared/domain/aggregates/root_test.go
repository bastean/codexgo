package aggregates_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type RootTestSuite struct {
	suite.Default
}

func (s *RootTestSuite) TestCreationStampErrOverwriteExisting() {
	root := aggregates.Mother().RootValid()

	s.NoError(root.CreationStamp())

	err := root.CreationStamp()

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "CreationStamp",
		What:  "Cannot overwrite an existing stamp",
	}}

	s.Equal(expected, actual)
}

func TestUnitRootSuite(t *testing.T) {
	suite.Run(t, new(RootTestSuite))
}
