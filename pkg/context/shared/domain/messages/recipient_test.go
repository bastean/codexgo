package messages_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type RecipientTestSuite struct {
	suite.Suite
}

func (s *RecipientTestSuite) TestWithValidValue() {
	components := &messages.RecipientComponents{
		Service: "user",
		Entity:  "user",
		Action:  "send confirmation",
		Event:   "created",
		Status:  messages.Status.Succeeded,
	}

	expected := messages.Recipient("user.user.send_confirmation_on_created_succeeded")

	actual := messages.NewRecipient(components)

	s.Equal(expected, actual)
}

func (s *RecipientTestSuite) TestWithInvalidValue() {
	s.Panics(func() { messages.NewRecipient(new(messages.RecipientComponents)) })
}

func TestUnitRecipientSuite(t *testing.T) {
	suite.Run(t, new(RecipientTestSuite))
}
