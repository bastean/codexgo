package user_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/stretchr/testify/suite"
)

type UsernameValueObjectTestSuite struct {
	suite.Suite
}

func (suite *UsernameValueObjectTestSuite) SetupTest() {}

func (suite *UsernameValueObjectTestSuite) TestWithInvalidLength() {
	value, err := user.UsernameWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewUsername",
		What:  "username must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Username": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *UsernameValueObjectTestSuite) TestWithInvalidAlphanumeric() {
	value, err := user.UsernameWithInvalidAlphanumeric()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewUsername",
		What:  "username must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Username": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitUsernameValueObjectSuite(t *testing.T) {
	suite.Run(t, new(UsernameValueObjectTestSuite))
}
