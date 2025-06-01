package password_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/notification/application/password"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
)

type PasswordTestSuite struct {
	suite.Frozen
	SUT      roles.EventConsumer
	password *password.Case
	transfer *transport.TransferMock
}

func (s *PasswordTestSuite) SetupSuite() {
	s.transfer = new(transport.TransferMock)

	s.password = &password.Case{
		Transfer: s.transfer,
	}

	s.SUT = &password.Consumer{
		Case: s.password,
	}
}

func (s *PasswordTestSuite) TestConsumer() {
	attributes := password.Mother().EventAttributesValid()

	aggregate, err := recipient.New(&recipient.Required{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
	})

	s.NoError(err)

	aggregate.ResetToken = values.Mother().IDNew(attributes.ResetToken)

	event := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.transfer.Mock.On("Submit", aggregate)

	s.NoError(s.SUT.On(event))

	s.transfer.Mock.AssertExpectations(s.T())
}

func TestUnitPasswordSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
