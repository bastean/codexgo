package user_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type PlainPasswordTestSuite struct {
	suite.Default
}

func (s *PlainPasswordTestSuite) TestWithInvalidLength() {
	value, err := user.Mother().PlainPasswordInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "user/*PlainPassword/Validate",
		What:  "Password must be between 8 to 64 characters",
		Why: errors.Meta{
			"Password": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitPlainPasswordSuite(t *testing.T) {
	suite.Run(t, new(PlainPasswordTestSuite))
}

type PasswordTestSuite struct {
	suite.Default
}

func (s *PasswordTestSuite) TestWithInvalidValue() {
	_, err := user.Mother().PasswordInvalid()

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "user/*Password/Validate",
		What:  "Password is required",
	}}

	s.Equal(expected, actual)
}

func TestUnitPasswordSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
