package valueobj_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/stretchr/testify/suite"
)

type PasswordValueObjectTestSuite struct {
	suite.Suite
}

func (suite *PasswordValueObjectTestSuite) SetupTest() {}

func (suite *PasswordValueObjectTestSuite) TestPassword() {
	password, err := valueobj.WithInvalidPasswordLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewPassword",
		What:  "password must be between " + "8" + " to " + "64" + " characters",
		Why: errors.Meta{
			"Password": password,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitPasswordValueObjectSuite(t *testing.T) {
	suite.Run(t, new(PasswordValueObjectTestSuite))
}
