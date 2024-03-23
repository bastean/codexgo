package sharedValueObject_test

import (
	"testing"

	sharedValueObjectMother "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type IdValueObjectTestSuite struct {
	suite.Suite
}

func (suite *IdValueObjectTestSuite) SetupTest() {}

func (suite *IdValueObjectTestSuite) TestId() {
	msg := "Id Invalid"

	suite.PanicsWithError(msg, func() { sharedValueObjectMother.InvalidId() })
}

func TestIdValueObjectSuite(t *testing.T) {
	suite.Run(t, new(IdValueObjectTestSuite))
}
