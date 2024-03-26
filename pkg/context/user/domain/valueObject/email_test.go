package valueObject_test

import (
	"testing"

	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type EmailValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EmailValueObjectTestSuite) SetupTest() {}

func (suite *EmailValueObjectTestSuite) TestEmail() {
	msg := "Email Invalid"

	suite.PanicsWithError(msg, func() { valueObjectMother.InvalidEmail() })
}

func TestEmailValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EmailValueObjectTestSuite))
}
