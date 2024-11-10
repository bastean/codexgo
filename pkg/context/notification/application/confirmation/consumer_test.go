package confirmation_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events/user"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/transports"
)

type ConfirmationTestSuite struct {
	suite.Suite
	sut          events.Consumer
	confirmation cases.Confirmation
	transfer     *transports.TransferMock[*user.CreatedSucceededAttributes]
}

func (suite *ConfirmationTestSuite) SetupTest() {
	suite.transfer = new(transports.TransferMock[*user.CreatedSucceededAttributes])

	suite.confirmation = &confirmation.Confirmation{
		Transfer: suite.transfer,
	}

	suite.sut = &confirmation.Consumer{
		Confirmation: suite.confirmation,
	}
}

func (suite *ConfirmationTestSuite) TestConsumer() {
	event := messages.RandomWithAttributes[events.Event](new(user.CreatedSucceededAttributes), true)

	suite.transfer.On("Submit", event.Attributes)

	suite.NoError(suite.sut.On(event))

	suite.transfer.AssertExpectations(suite.T())
}

func TestUnitNotificationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
