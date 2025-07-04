package messages_test

import (
	"fmt"
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type RecipientTestSuite struct {
	suite.Default
}

func (s *RecipientTestSuite) TestSentinel() {
	s.Equal(`^([a-z0-9]{1,30})\.([a-z]{1,30})\.([a-z_]{1,30})_on_([a-z]{1,30})_(queued|succeeded|failed|done)$`, messages.RExRecipient)
}

func (s *RecipientTestSuite) TestWithValidValue() {
	components := messages.Mother().RecipientComponentsValid()

	actual := messages.Mother().RecipientValidWithComponents(components).Value()

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
	components := messages.Mother().RecipientComponentsInvalid()

	expected := fmt.Sprintf("(messages/*Recipient/Validate): Recipient has an invalid nomenclature %q", messages.FormatRecipient(components))

	s.PanicsWithValue(expected, func() { messages.Mother().RecipientInvalidWithComponents(components) })
}

func (s *RecipientTestSuite) TestParseRecipient() {
	expected := messages.Mother().RecipientComponentsValid()

	var actual *messages.RecipientComponents

	s.NotPanics(func() { actual = messages.ParseRecipient(messages.FormatRecipient(expected)) })

	s.Equal(expected, actual)
}

func TestUnitRecipientSuite(t *testing.T) {
	suite.Run(t, new(RecipientTestSuite))
}
