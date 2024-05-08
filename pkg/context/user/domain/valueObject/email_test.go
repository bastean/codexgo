package valueObject_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	"github.com/stretchr/testify/suite"
)

type EmailValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EmailValueObjectTestSuite) SetupTest() {}

func (suite *EmailValueObjectTestSuite) TestEmail() {
	// TODO!(test): fix unit tests

	invalidEmail := ""

	_, err := valueObject.NewEmail(invalidEmail)

	suite.Error(err)

	/*
		expected := errs.NewInvalidValueError(&errs.Bubble{
				Where: "NewEmail",
				What:  "invalid format",
				Why:   errs.Meta{
					"Email": invalidEmail,
				},
			})
	*/

}

func TestUnitEmailValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EmailValueObjectTestSuite))
}
