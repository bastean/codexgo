package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type ResponseTestSuite struct {
	suite.Suite
}

func (suite *ResponseTestSuite) TestWithInvalidLength() {
	value, err := components.ResponseWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewResponse",
		What:  "Response must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Response": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *ResponseTestSuite) TestWithInvalidAlpha() {
	value, err := components.ResponseWithInvalidAlpha()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewResponse",
		What:  "Response must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Response": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitResponseSuite(t *testing.T) {
	suite.Run(t, new(ResponseTestSuite))
}
