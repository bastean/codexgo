package valueObject_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type PasswordValueObjectTestSuite struct {
	suite.Suite
}

func (suite *PasswordValueObjectTestSuite) SetupTest() {}

func (suite *PasswordValueObjectTestSuite) TestPassword() {
	password, err := valueObjectMother.WithInvalidPasswordLength()

	expected := errs.NewInvalidValueError(&errs.Bubble{
		Where: "NewPassword",
		What:  "must be between " + "8" + " to " + "64" + " characters",
		Why: errs.Meta{
			"Password": password,
		},
	})

	var actual *errs.InvalidValueError

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func TestUnitPasswordValueObjectSuite(t *testing.T) {
	suite.Run(t, new(PasswordValueObjectTestSuite))
}
