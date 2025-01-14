package confirmation_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type ConfirmationTestSuite struct {
	suite.Suite
	SUT          roles.EventConsumer
	confirmation cases.Confirmation
	transfer     *transport.TransferMock[*events.UserCreatedSucceededAttributes]
}

func (s *ConfirmationTestSuite) SetupTest() {
	s.transfer = new(transport.TransferMock[*events.UserCreatedSucceededAttributes])

	s.confirmation = &confirmation.Case{
		Transfer: s.transfer,
	}

	s.SUT = &confirmation.Consumer{
		Confirmation: s.confirmation,
	}
}

func (s *ConfirmationTestSuite) TestConsumer() {
	event := messages.RandomWithAttributes(new(events.UserCreatedSucceededAttributes), true)

	s.transfer.Mock.On("Submit", event.Attributes)

	s.NoError(s.SUT.On(event))

	s.transfer.Mock.AssertExpectations(s.T())
}

func TestUnitConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
