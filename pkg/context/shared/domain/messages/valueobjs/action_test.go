package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
	"github.com/stretchr/testify/suite"
)

type ActionValueObjectTestSuite struct {
	suite.Suite
}

func (suite *ActionValueObjectTestSuite) SetupTest() {}

func (suite *ActionValueObjectTestSuite) TestWithInvalidLength() {
	value, err := valueobjs.ActionWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewAction",
		What:  "Action must be between 1 to 20 characters",
		Why: errors.Meta{
			"Action": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitActionValueObjectSuite(t *testing.T) {
	suite.Run(t, new(ActionValueObjectTestSuite))
}
