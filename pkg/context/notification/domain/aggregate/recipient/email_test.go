package recipient_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type EmailTestSuite struct {
	suite.Default
}

func (s *EmailTestSuite) TestWithInvalidValue() {
	value, err := recipient.Mother().EmailInvalid()

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
