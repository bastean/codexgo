package user_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type PasswordTestSuite struct {
	suite.Suite
}

func (suite *PasswordTestSuite) SetupTest() {}

func (suite *PasswordTestSuite) TestWithInvalidLength() {
	value, err := user.PasswordWithInvalidLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewPassword",
		What:  "Password must be between 8 to 64 characters",
		Why: errors.Meta{
			"Password": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitPasswordSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
