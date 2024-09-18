package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type EntityTestSuite struct {
	suite.Suite
}

func (suite *EntityTestSuite) SetupTest() {}

func (suite *EntityTestSuite) TestWithInvalidLength() {
	value, err := components.EntityWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEntity",
		What:  "Entity must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Entity": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *EntityTestSuite) TestWithInvalidAlpha() {
	value, err := components.EntityWithInvalidAlpha()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEntity",
		What:  "Entity must be between 1 to 20 characters and be alpha only",
		Why: errors.Meta{
			"Entity": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEntitySuite(t *testing.T) {
	suite.Run(t, new(EntityTestSuite))
}
