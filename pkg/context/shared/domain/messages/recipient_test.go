package messages_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/stretchr/testify/suite"
)

type RecipientNameTestSuite struct {
	suite.Suite
}

func (suite *RecipientNameTestSuite) SetupTest() {}

func (suite *RecipientNameTestSuite) TestWithValidValue() {
	components := &messages.RecipientNameComponents{
		Service: "user",
		Entity:  "user",
		Action:  "send confirmation",
		Event:   "created",
		Status:  messages.Status.Succeeded,
	}

	expected := "user.user.send_confirmation_on_created_succeeded"

	actual := messages.NewRecipientName(components)

	suite.Equal(expected, actual)
}

func (suite *RecipientNameTestSuite) TestWithInvalidValue() {
	components := &messages.RecipientNameComponents{}
	suite.Panics(func() { messages.NewRecipientName(components) })
}

func TestUnitRecipientNameSuite(t *testing.T) {
	suite.Run(t, new(RecipientNameTestSuite))
}
