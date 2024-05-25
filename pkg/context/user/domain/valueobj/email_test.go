package valueobj_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/stretchr/testify/suite"
)

type EmailValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EmailValueObjectTestSuite) SetupTest() {}

func (suite *EmailValueObjectTestSuite) TestEmail() {
	email, err := valueobj.InvalidEmail()

	var actual *serror.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := serror.InvalidValue{Bubble: &serror.Bubble{
		When:  actual.When,
		Where: "NewEmail",
		What:  "invalid email format",
		Why: serror.Meta{
			"Email": email,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEmailValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EmailValueObjectTestSuite))
}
