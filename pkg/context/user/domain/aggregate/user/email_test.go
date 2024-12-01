package user_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type EmailTestSuite struct {
	suite.Suite
}

func (s *EmailTestSuite) TestWithInvalidValue() {
	value, err := user.EmailWithInvalidValue()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEmail",
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
