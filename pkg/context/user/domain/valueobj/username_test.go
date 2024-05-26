package valueobj_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/stretchr/testify/suite"
)

type UsernameValueObjectTestSuite struct {
	suite.Suite
}

func (suite *UsernameValueObjectTestSuite) SetupTest() {}

func (suite *UsernameValueObjectTestSuite) TestUsernameWithInvalidLength() {
	username, err := valueobj.WithInvalidUsernameLength()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewUsername",
		What:  "username must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Username": username,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *UsernameValueObjectTestSuite) TestUsernameWithInvalidAlphanumeric() {
	username, err := valueobj.WithInvalidUsernameAlphanumeric()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewUsername",
		What:  "username must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Username": username,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitUsernameValueObjectSuite(t *testing.T) {
	suite.Run(t, new(UsernameValueObjectTestSuite))
}
