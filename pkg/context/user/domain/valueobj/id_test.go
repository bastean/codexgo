package valueobj_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/serror"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/stretchr/testify/suite"
)

type IdValueObjectTestSuite struct {
	suite.Suite
}

func (suite *IdValueObjectTestSuite) SetupTest() {}

func (suite *IdValueObjectTestSuite) TestId() {
	id, err := valueobj.InvalidId()

	expected := serror.NewInvalidValue(&serror.Bubble{
		Where: "NewId",
		What:  "invalid id format",
		Why: serror.Meta{
			"Id": id,
		},
	})

	var actual *serror.InvalidValue

	suite.ErrorAs(err, &actual)

	suite.EqualError(expected, actual.Error())
}

func TestUnitIdValueObjectSuite(t *testing.T) {
	suite.Run(t, new(IdValueObjectTestSuite))
}
