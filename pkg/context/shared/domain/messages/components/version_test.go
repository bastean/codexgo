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

func (suite *VersionTestSuite) TestWithInvalidValue() {
	value, err := components.VersionWithInvalidValue()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewVersion",
		What:  "Version must be numeric only",
		Why: errors.Meta{
			"Version": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitVersionSuite(t *testing.T) {
	suite.Run(t, new(VersionTestSuite))
}
