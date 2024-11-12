package user_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/stretchr/testify/suite"
)

type PasswordTestSuite struct {
	suite.Suite
}

func (s *PasswordTestSuite) TestWithInvalidLength() {
	value, err := user.PasswordWithInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewPassword",
		What:  "Password must be between 8 to 64 characters",
		Why: errors.Meta{
			"Password": value,
		},
	}}

	s.EqualError(expected, actual.Error())
}

func TestUnitPasswordSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
