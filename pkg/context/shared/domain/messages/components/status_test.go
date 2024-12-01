package components_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages/components"
)

type StatusTestSuite struct {
	suite.Suite
}

func (s *StatusTestSuite) TestWithInvalidValue() {
	value, err := components.StatusWithInvalidValue()

	var actual *errors.InvalidValue

	s.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewStatus",
		What:  "Status must be only one of these values: Queued, Succeeded, Failed, Done",
		Why: errors.Meta{
			"Status": value,
		},
	}}

	s.Equal(expected, actual)
}

func TestUnitStatusSuite(t *testing.T) {
	suite.Run(t, new(StatusTestSuite))
}
