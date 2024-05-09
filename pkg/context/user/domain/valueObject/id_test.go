package valueObject_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/errs"
	valueObjectMother "github.com/bastean/codexgo/pkg/context/user/domain/valueObject/mother"
	"github.com/stretchr/testify/suite"
)

type IdValueObjectTestSuite struct {
	suite.Suite
}

func (suite *IdValueObjectTestSuite) SetupTest() {}

func (suite *IdValueObjectTestSuite) TestId() {
	id, err := valueObjectMother.InvalidId()

	expected := errs.NewInvalidValueError(&errs.Bubble{
		Where: "NewId",
		What:  "invalid format",
		Why: errs.Meta{
			"Id": id,
		},
	})

	var actual *errs.InvalidValueError

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func TestUnitIdValueObjectSuite(t *testing.T) {
	suite.Run(t, new(IdValueObjectTestSuite))
}
