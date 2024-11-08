package messages_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
)

type RecipientTestSuite struct {
	suite.Suite
}

func (suite *RecipientTestSuite) TestWithValidValue() {
	components := &messages.RecipientComponents{
		Service: "user",
		Entity:  "user",
		Action:  "send confirmation",
		Event:   "created",
		Status:  messages.Status.Succeeded,
	}

	expected := events.Recipient("user.user.send_confirmation_on_created_succeeded")

	actual := messages.NewRecipient(components)

	suite.Equal(expected, actual)
}

func (suite *RecipientTestSuite) TestWithInvalidValue() {
	suite.Panics(func() { messages.NewRecipient(&messages.RecipientComponents{}) })
}

func TestUnitRecipientSuite(t *testing.T) {
	suite.Run(t, new(RecipientTestSuite))
}
