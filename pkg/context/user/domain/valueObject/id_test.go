package valueObject_test

import (
	"testing"

	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type IdValueObjectTestSuite struct {
	suite.Suite
}

func (suite *IdValueObjectTestSuite) SetupTest() {}

func (suite *IdValueObjectTestSuite) TestId() {
	msg := "Id Invalid"

	suite.PanicsWithError(msg, func() { valueObjectMother.InvalidId() })
}

func TestIdValueObjectSuite(t *testing.T) {
	suite.Run(t, new(IdValueObjectTestSuite))
}
