package valueObject_test

import (
	"testing"

	valueObjectMother "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
	"github.com/stretchr/testify/suite"
)

type UserValueObjectTestSuite struct {
	suite.Suite
}

func (suite *UserValueObjectTestSuite) SetupTest() {}

func (suite *UserValueObjectTestSuite) TestId() {
	msg := "Id Invalid"

	suite.PanicsWithError(msg, func() { valueObjectMother.InvalidId() })
}

func (suite *UserValueObjectTestSuite) TestEmail() {
	msg := "Email Invalid"

	suite.PanicsWithError(msg, func() { valueObjectMother.InvalidEmail() })
}

func (suite *UserValueObjectTestSuite) TestUsername() {
	msg := "Username must be between " + "2" + " to " + "20" + " characters and be alphanumeric only"

	suite.PanicsWithError(msg, func() { valueObjectMother.WithInvalidUsernameLength() })
	suite.PanicsWithError(msg, func() { valueObjectMother.WithInvalidUsernameAlphanumeric() })
}

func (suite *UserValueObjectTestSuite) TestPassword() {
	msg := "Password must be between " + "8" + " to " + "64" + " characters"

	suite.PanicsWithError(msg, func() { valueObjectMother.WithInvalidPasswordLength() })
}

func TestUserValueObjectSuite(t *testing.T) {
	suite.Run(t, new(UserValueObjectTestSuite))
}
