package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
	"github.com/stretchr/testify/suite"
)

type StatusValueObjectTestSuite struct {
	suite.Suite
}

func (suite *StatusValueObjectTestSuite) SetupTest() {}

func (suite *StatusValueObjectTestSuite) TestWithInvalidValue() {
	value, err := valueobjs.StatusWithInvalidValue()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewStatus",
		What:  "Status must be only one of these values: Queued, Succeeded, Failed, Done",
		Why: errors.Meta{
			"Status": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitStatusValueObjectSuite(t *testing.T) {
	suite.Run(t, new(StatusValueObjectTestSuite))
}
