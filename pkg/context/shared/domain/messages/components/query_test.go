package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type QueryTestSuite struct {
	suite.Suite
}

func (suite *QueryTestSuite) TestWithInvalidLength() {
	value, err := components.QueryWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewQuery",
		What:  "Query must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Query": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *QueryTestSuite) TestWithInvalidAlpha() {
	value, err := components.QueryWithInvalidAlpha()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewQuery",
		What:  "Query must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Query": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitQuerySuite(t *testing.T) {
	suite.Run(t, new(QueryTestSuite))
}
