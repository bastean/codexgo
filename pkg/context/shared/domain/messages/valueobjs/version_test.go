package valueobjs_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/valueobjs"
	"github.com/stretchr/testify/suite"
)

type VersionValueObjectTestSuite struct {
	suite.Suite
}

func (suite *VersionValueObjectTestSuite) SetupTest() {}

func (suite *VersionValueObjectTestSuite) TestWithInvalidValue() {
	value, err := valueobjs.VersionWithInvalidValue()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewVersion",
		What:  "Version must be numeric only",
		Why: errors.Meta{
			"Version": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitVersionValueObjectSuite(t *testing.T) {
	suite.Run(t, new(VersionValueObjectTestSuite))
}
