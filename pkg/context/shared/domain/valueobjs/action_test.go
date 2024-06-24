package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
	"github.com/stretchr/testify/suite"
)

type ActionValueObjectTestSuite struct {
	suite.Suite
}

func (suite *ActionValueObjectTestSuite) SetupTest() {}

func (suite *ActionValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.ActionWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewAction",
		What:  "action must be between " + "1" + " to " + "20" + " characters",
		Why: errors.Meta{
			"Action": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitActionValueObjectSuite(t *testing.T) {
	suite.Run(t, new(ActionValueObjectTestSuite))
}
