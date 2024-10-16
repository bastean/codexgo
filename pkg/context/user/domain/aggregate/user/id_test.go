package user_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type IdTestSuite struct {
	suite.Suite
}

func (suite *IdTestSuite) TestWithInvalidValue() {
	value, err := user.IdWithInvalidValue()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewId",
		What:  "Invalid UUID4 format",
		Why: errors.Meta{
			"Id": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitIdSuite(t *testing.T) {
	suite.Run(t, new(IdTestSuite))
}
