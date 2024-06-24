package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/shared/domain/valueobjs"
	"github.com/stretchr/testify/suite"
)

type StatusValueObjectTestSuite struct {
	suite.Suite
}

func (suite *StatusValueObjectTestSuite) SetupTest() {}

func (suite *StatusValueObjectTestSuite) TestWithInvalidValue() {
	value, err := valueobjs.StatusWithInvalidValue()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewStatus",
		What:  "status must be only one of these values: queued, succeeded, failed, done",
		Why: errors.Meta{
			"Status": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitStatusValueObjectSuite(t *testing.T) {
	suite.Run(t, new(StatusValueObjectTestSuite))
}
