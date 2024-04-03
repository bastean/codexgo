package valueObject_test

import (
	"testing"

	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type PasswordValueObjectTestSuite struct {
	suite.Suite
}

func (suite *PasswordValueObjectTestSuite) SetupTest() {}

func (suite *PasswordValueObjectTestSuite) TestPassword() {
	msg := "Password must be between " + "8" + " to " + "64" + " characters"

	suite.PanicsWithError(msg, func() { valueObjectMother.WithInvalidPasswordLength() })
}

func TestUnitPasswordValueObjectSuite(t *testing.T) {
	suite.Run(t, new(PasswordValueObjectTestSuite))
}
