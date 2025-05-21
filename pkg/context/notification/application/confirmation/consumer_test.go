package confirmation_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/notification/application/confirmation"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type ConfirmationTestSuite struct {
	suite.Frozen
	SUT          roles.EventConsumer
	confirmation *confirmation.Case
	transfer     *transport.TransferMock
}

func (s *ConfirmationTestSuite) SetupSuite() {
	s.transfer = new(transport.TransferMock)

	s.confirmation = &confirmation.Case{
		Transfer: s.transfer,
	}

	s.SUT = &confirmation.Consumer{
		Case: s.confirmation,
	}
}

func (s *ConfirmationTestSuite) TestConsumer() {
	attributes := confirmation.Mother().EventAttributesValid()

	aggregate, err := recipient.New(&recipient.Required{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
	})

	s.NoError(err)

	aggregate.VerifyToken = recipient.Mother().IDNew(attributes.VerifyToken)

	event := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.transfer.Mock.On("Submit", aggregate)

	s.NoError(s.SUT.On(event))

	s.transfer.Mock.AssertExpectations(s.T())
}

func TestUnitConfirmationSuite(t *testing.T) {
	suite.Run(t, new(ConfirmationTestSuite))
}
