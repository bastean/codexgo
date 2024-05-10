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

	expected := serror.NewInvalidValue(&serror.Bubble{
		Where: "NewEmail",
		What:  "invalid email format",
		Why: serror.Meta{
			"Email": email,
		},
	})

	var actual *serror.InvalidValue

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func TestUnitEmailValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EmailValueObjectTestSuite))
}
