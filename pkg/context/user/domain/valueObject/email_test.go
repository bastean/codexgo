package valueObject_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type EmailValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EmailValueObjectTestSuite) SetupTest() {}

func (suite *EmailValueObjectTestSuite) TestEmail() {
	email, err := valueObjectMother.InvalidEmail()

	expected := errs.NewInvalidValueError(&errs.Bubble{
		Where: "NewEmail",
		What:  "invalid format",
		Why: errs.Meta{
			"Email": email,
		},
	})

	var actual *errs.InvalidValueError

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func TestUnitEmailValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EmailValueObjectTestSuite))
}
