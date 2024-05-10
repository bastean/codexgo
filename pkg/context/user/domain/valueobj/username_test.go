package valueobj_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/stretchr/testify/suite"
)

type UsernameValueObjectTestSuite struct {
	suite.Suite
}

func (suite *UsernameValueObjectTestSuite) SetupTest() {}

func (suite *UsernameValueObjectTestSuite) TestUsernameWithInvalidLength() {
	username, err := valueobj.WithInvalidUsernameLength()

	expected := serror.NewInvalidValueError(&serror.Bubble{
		Where: "NewUsername",
		What:  "must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: serror.Meta{
			"Username": username,
		},
	})

	var actual *serror.InvalidValueError

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func (suite *UsernameValueObjectTestSuite) TestUsernameWithInvalidAlphanumeric() {
	username, err := valueobj.WithInvalidUsernameAlphanumeric()

	expected := serror.NewInvalidValueError(&serror.Bubble{
		Where: "NewUsername",
		What:  "must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: serror.Meta{
			"Username": username,
		},
	})

	var actual *serror.InvalidValueError

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func TestUnitUsernameValueObjectSuite(t *testing.T) {
	suite.Run(t, new(UsernameValueObjectTestSuite))
}
