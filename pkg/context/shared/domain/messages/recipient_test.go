package messages_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type RecipientTestSuite struct {
	suite.Suite
}

func (s *RecipientTestSuite) SetupSuite() {
	s.Equal(messages.RExRecipient, `^([a-z0-9]{1,30})\.([a-z]{1,30})\.([a-z_]{1,30})_on_([a-z]{1,30})_(queued|succeeded|failed|done)$`)
}

func (s *RecipientTestSuite) TestWithValidValue() {
	components := messages.Mother.RecipientComponentsValid()

	recipient, err := values.New[*messages.Recipient](messages.ParseRecipient(components))

	s.NoError(err)

	actual := recipient.Value()

	expected := fmt.Sprintf("%s.%s.%s_on_%s_%s",
		components.Service,
		components.Entity,
		components.Trigger,
		components.Action,
		components.Status,
	)

	s.Equal(expected, actual)
}

func (s *RecipientTestSuite) TestWithInvalidValue() {
	expected := "(Validate): Recipient has an invalid nomenclature"

	s.PanicsWithValue(expected, func() {
		_, _ = values.New[*messages.Recipient](messages.ParseRecipient(messages.Mother.RecipientComponentsInvalid()))
	})
}

func TestUnitRecipientSuite(t *testing.T) {
	suite.Run(t, new(RecipientTestSuite))
}
