package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type ActionTestSuite struct {
	suite.Suite
}

func (suite *ActionTestSuite) SetupTest() {}

func (suite *ActionTestSuite) TestWithInvalidLength() {
	value, err := components.ActionWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewAction",
		What:  "Action must be between 1 to 20 characters",
		Why: errors.Meta{
			"Action": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitActionSuite(t *testing.T) {
	suite.Run(t, new(ActionTestSuite))
}
