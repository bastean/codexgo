package user_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type UsernameTestSuite struct {
	suite.Suite
}

func (suite *UsernameTestSuite) TestWithInvalidLength() {
	value, err := user.UsernameWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewUsername",
		What:  "Username must be between 2 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Username": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *UsernameTestSuite) TestWithInvalidAlphanumeric() {
	value, err := user.UsernameWithInvalidAlphanumeric()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewUsername",
		What:  "Username must be between 2 to 20 characters and be alphanumeric only",
		Why: errors.Meta{
			"Username": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitUsernameSuite(t *testing.T) {
	suite.Run(t, new(UsernameTestSuite))
}
