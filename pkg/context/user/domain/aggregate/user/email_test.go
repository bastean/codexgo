package user_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type EmailTestSuite struct {
	suite.Default
}

func (s *EmailTestSuite) TestWithInvalidValue() {
	value, err := user.Mother().EmailInvalid()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Validate",
		What:  "Invalid email format",
		Why: errors.Meta{
			"Email": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitEmailSuite(t *testing.T) {
	suite.Run(t, new(EmailTestSuite))
}
