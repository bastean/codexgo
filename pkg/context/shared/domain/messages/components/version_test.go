package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type VersionTestSuite struct {
	suite.Suite
}

func (s *VersionTestSuite) TestWithInvalidValue() {
	value, err := components.VersionWithInvalidValue()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewVersion",
		What:  "Version must be numeric only",
		Why: errors.Meta{
			"Version": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitVersionSuite(t *testing.T) {
	suite.Run(t, new(VersionTestSuite))
}
