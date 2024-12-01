package user_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type PlainPasswordTestSuite struct {
	suite.Suite
}

type CipherPasswordTestSuite struct {
	suite.Suite
}

func (s *PlainPasswordTestSuite) TestWithInvalidLength() {
	value, err := user.PlainPasswordWithInvalidLength()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewPlainPassword",
		What:  "Password must be between 8 to 64 characters",
		Why: errors.Meta{
			"Password": value,
		},
	}}

	s.Equal(expected, actual)
}

func (s *CipherPasswordTestSuite) TestWithInvalidValue() {
	value, err := user.CipherPasswordWithInvalidValue()

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewCipherPassword",
		What:  "Cipher Password is required",
		Why: errors.Meta{
			"Password": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitPlainPasswordSuite(t *testing.T) {
	suite.Run(t, new(PlainPasswordTestSuite))
}

func TestUnitCipherPasswordSuite(t *testing.T) {
	suite.Run(t, new(CipherPasswordTestSuite))
}
