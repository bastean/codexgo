package root_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/root"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type RootTestSuite struct {
	suite.Default
}

func (s *RootTestSuite) TestCreationStampErrOverwriteExisting() {
	aggregate := root.Mother().RootValid()

	s.NoError(aggregate.CreationStamp())

	err := aggregate.CreationStamp()

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "root/*Root/CreationStamp",
		What:  "Cannot overwrite an existing stamp",
	}}

	s.Equal(expected, actual)
}

func TestUnitRootSuite(t *testing.T) {
	suite.Run(t, new(RootTestSuite))
}
