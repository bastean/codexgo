package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
	"github.com/stretchr/testify/suite"
)

type TypeValueObjectTestSuite struct {
	suite.Suite
}

func (suite *TypeValueObjectTestSuite) SetupTest() {}

func (suite *TypeValueObjectTestSuite) TestWithInvalidValue() {
	value, err := valueobjs.TypeWithInvalidValue()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewType",
		What:  "Type must be only one of these values: Event, Command",
		Why: errors.Meta{
			"Type": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitTypeValueObjectSuite(t *testing.T) {
	suite.Run(t, new(TypeValueObjectTestSuite))
}
