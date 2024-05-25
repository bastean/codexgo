package valueobj_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/stretchr/testify/suite"
)

type PasswordValueObjectTestSuite struct {
	suite.Suite
}

func (suite *PasswordValueObjectTestSuite) SetupTest() {}

func (suite *PasswordValueObjectTestSuite) TestPassword() {
	password, err := valueobj.WithInvalidPasswordLength()

	var actual *serror.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := serror.InvalidValue{Bubble: &serror.Bubble{
		When:  actual.When,
		Where: "NewPassword",
		What:  "password must be between " + "8" + " to " + "64" + " characters",
		Why: serror.Meta{
			"Password": password,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitPasswordValueObjectSuite(t *testing.T) {
	suite.Run(t, new(PasswordValueObjectTestSuite))
}