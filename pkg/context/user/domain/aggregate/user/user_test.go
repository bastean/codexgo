package user_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type UserTestSuite struct {
	suite.Suite
}

func (s *UserTestSuite) TestValidateVerifyErrDoNotMatch() {
	aggregate := user.Random()

	token := user.IDWithValidValue()

	err := aggregate.ValidateVerify(token)

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "ValidateVerify",
		What:  "Tokens do not match",
		Why: errors.Meta{
			"Received": token,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
