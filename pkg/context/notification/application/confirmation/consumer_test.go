package confirmation_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports"
	"github.com/stretchr/testify/suite"
)

type ConfirmationTestSuite struct {
	suite.Suite
	sut          events.Consumer
	confirmation cases.Confirmation
	transfer     *transports.TransferMock[*user.CreatedSucceededAttributes]
}

func (s *ConfirmationTestSuite) SetupTest() {
	s.transfer = new(transports.TransferMock[*user.CreatedSucceededAttributes])

	s.confirmation = &confirmation.Case{
		Transfer: s.transfer,
	}

	s.sut = &confirmation.Consumer{
		Confirmation: s.confirmation,
	}
}

func (s *ConfirmationTestSuite) TestConsumer() {
	event := messages.RandomWithAttributes[events.Event](new(user.CreatedSucceededAttributes), true)

	s.transfer.On("Submit", event.Attributes)

	s.NoError(s.sut.On(event))

	s.transfer.AssertExpectations(s.T())
}

func TestUnitConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
