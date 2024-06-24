package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
	"github.com/stretchr/testify/suite"
)

type TypeValueObjectTestSuite struct {
	suite.Suite
}

func (suite *TypeValueObjectTestSuite) SetupTest() {}

func (suite *TypeValueObjectTestSuite) TestWithInvalidValue() {
	value, err := valueobjs.TypeWithInvalidValue()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewType",
		What:  "type must be only one of these values: event, command",
		Why: errors.Meta{
			"Type": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitTypeValueObjectSuite(t *testing.T) {
	suite.Run(t, new(TypeValueObjectTestSuite))
}
