package valueObject_test

import (
	"testing"

	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type UsernameValueObjectTestSuite struct {
	suite.Suite
}

func (suite *UsernameValueObjectTestSuite) SetupTest() {}

func (suite *UsernameValueObjectTestSuite) TestUsername() {
	msg := "Username must be between " + "2" + " to " + "20" + " characters and be alphanumeric only"

	suite.PanicsWithError(msg, func() { valueObjectMother.WithInvalidUsernameLength() })
	suite.PanicsWithError(msg, func() { valueObjectMother.WithInvalidUsernameAlphanumeric() })
}

func TestUsernameValueObjectSuite(t *testing.T) {
	suite.Run(t, new(UsernameValueObjectTestSuite))
}
