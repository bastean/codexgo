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

func (suite *EmailTestSuite) SetupTest() {}

func (suite *EmailTestSuite) TestWithInvalidValue() {
	value, err := user.EmailWithInvalidValue()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEmail",
		What:  "Invalid email format",
		Why: errors.Meta{
			"Email": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEmailSuite(t *testing.T) {
	suite.Run(t, new(EmailTestSuite))
}
