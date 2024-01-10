package valueObjects_test

import (
	"testing"

	create "github.com/bastean/codexgo/test/contexts/crud/user/domain/valueObjects"
	"github.com/stretchr/testify/suite"
)

type UserValueObjectsTestSuite struct {
	suite.Suite
}

func (suite *UserValueObjectsTestSuite) SetupTest() {}

func (suite *UserValueObjectsTestSuite) TestId() {
	msg := "Id Invalid"

	suite.PanicsWithError(msg, func() { create.InvalidId() })
}

func (suite *UserValueObjectsTestSuite) TestEmail() {
	msg := "Email Invalid"

	suite.PanicsWithError(msg, func() { create.InvalidEmail() })
}

func (suite *UserValueObjectsTestSuite) TestUsername() {
	msg := "Username must be between " + "2" + " to " + "20" + " characters and be alphanumeric only"

	suite.PanicsWithError(msg, func() { create.WithInvalidUsernameLength() })
	suite.PanicsWithError(msg, func() { create.WithInvalidUsernameAlphanumeric() })
}

func (suite *UserValueObjectsTestSuite) TestPassword() {
	msg := "Password must be between " + "8" + " to " + "64" + " characters"

	suite.PanicsWithError(msg, func() { create.WithInvalidPasswordLength() })
}

func TestUserValueObjectsTestSuite(t *testing.T) {
	suite.Run(t, new(UserValueObjectsTestSuite))
}
