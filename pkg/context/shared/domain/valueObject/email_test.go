package sharedValueObject_test

import (
	"testing"

	sharedValueObjectMother "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type EmailValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EmailValueObjectTestSuite) SetupTest() {}

func (suite *EmailValueObjectTestSuite) TestEmail() {
	msg := "Email Invalid"

	suite.PanicsWithError(msg, func() { sharedValueObjectMother.InvalidEmail() })
}

func TestEmailValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EmailValueObjectTestSuite))
}
