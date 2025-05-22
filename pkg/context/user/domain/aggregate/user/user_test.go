package user_test

import (
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type UserTestSuite struct {
	suite.Default
}

func (s *UserTestSuite) TestValidateVerifyTokenErrDoNotMatch() {
	aggregate := user.Mother().UserValidFromPrimitive()

	token := user.Mother().IDValid()

	err := aggregate.ValidateVerifyToken(token)

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "ValidateVerifyToken",
		What:  "Tokens do not match",
		Why: errors.Meta{
			"Received": token.Value(),
		},
	}}

	s.Equal(expected, actual)
}

func (s *UserTestSuite) TestValidateResetTokenErrDoNotMatch() {
	aggregate := user.Mother().UserValidFromPrimitive()

	token := user.Mother().IDValid()

	err := aggregate.ValidateResetToken(token)

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "ValidateResetToken",
		What:  "Tokens do not match",
		Why: errors.Meta{
			"Received": token.Value(),
		},
	}}

	s.Equal(expected, actual)
}

func (s *UserTestSuite) TestCreatedSucceededKey() {
	actual := user.CreatedSucceededKey.Value()

	expected := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s",
		"codexgo",
		"user",
		"1",
		"event",
		"user",
		"created",
		"succeeded",
	)

	s.Equal(expected, actual)
}

func (s *UserTestSuite) TestResetQueuedKey() {
	actual := user.ResetQueuedKey.Value()

	expected := fmt.Sprintf("%s.%s.%s.%s.%s.%s.%s",
		"codexgo",
		"user",
		"1",
		"event",
		"user",
		"reset",
		"queued",
	)

	s.Equal(expected, actual)
}

func TestUnitUserSuite(t *testing.T) {
	suite.Run(t, new(UserTestSuite))
}
