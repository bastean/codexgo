package valueObject_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type UsernameValueObjectTestSuite struct {
	suite.Suite
}

func (suite *UsernameValueObjectTestSuite) SetupTest() {}

func (suite *UsernameValueObjectTestSuite) TestUsernameWithInvalidLength() {
	username, err := valueObjectMother.WithInvalidUsernameLength()

	expected := errs.NewInvalidValueError(&errs.Bubble{
		Where: "NewUsername",
		What:  "must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: errs.Meta{
			"Username": username,
		},
	})

	var actual *errs.InvalidValueError

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func (suite *UsernameValueObjectTestSuite) TestUsernameWithInvalidAlphanumeric() {
	username, err := valueObjectMother.WithInvalidUsernameAlphanumeric()

	expected := errs.NewInvalidValueError(&errs.Bubble{
		Where: "NewUsername",
		What:  "must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: errs.Meta{
			"Username": username,
		},
	})

	var actual *errs.InvalidValueError

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func TestUnitUsernameValueObjectSuite(t *testing.T) {
	suite.Run(t, new(UsernameValueObjectTestSuite))
}
