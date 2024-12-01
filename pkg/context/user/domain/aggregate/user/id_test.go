package user_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
)

type IDTestSuite struct {
	suite.Suite
}

func (s *IDTestSuite) TestWithInvalidValue() {
	value, err := user.IDWithInvalidValue()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewID",
		What:  "Invalid UUID4 format",
		Why: errors.Meta{
			"ID": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitIDSuite(t *testing.T) {
	suite.Run(t, new(IDTestSuite))
}
