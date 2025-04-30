package messages_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type RecipientTestSuite struct {
	suite.Suite
}

func (s *RecipientTestSuite) SetupSuite() {
	s.Equal(messages.RExRecipientTrigger, `([a-z_]{1,20})`)

	s.Equal(messages.RExRecipientComponents, `^([a-z0-9]{1,20})\.([a-z]{1,20})\.([a-z_]{1,20})_on_([a-z]{1,20})_(queued|succeeded|failed|done)$`)
}

func (s *RecipientTestSuite) TestWithValidValue() {
	recipient, err := values.New[*messages.Recipient](messages.ParseRecipient(&messages.RecipientComponents{
		Service: "user",
		Entity:  "user",
		Trigger: "send_confirmation",
		Action:  "created",
		Status:  messages.Status.Succeeded,
	}))

	s.NoError(err)

	actual := recipient.Value()

	expected := "user.user.send_confirmation_on_created_succeeded"

	s.Equal(expected, actual)
}

func (s *RecipientTestSuite) TestWithInvalidValue() {
	expected := "(Validate): Recipient has an invalid nomenclature"

	s.PanicsWithValue(expected, func() {
		_, _ = values.New[*messages.Recipient](messages.ParseRecipient(new(messages.RecipientComponents)))
	})
}

func TestUnitRecipientSuite(t *testing.T) {
	suite.Run(t, new(RecipientTestSuite))
}
