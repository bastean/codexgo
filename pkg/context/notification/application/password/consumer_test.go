package password_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/notification/application/password"
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/notification/infrastructure/transport"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type PasswordTestSuite struct {
	suite.Suite
	SUT      roles.EventConsumer
	password cases.Password
	transfer *transport.TransferMock[*events.UserResetQueuedAttributes]
}

func (s *PasswordTestSuite) SetupTest() {
	s.transfer = new(transport.TransferMock[*events.UserResetQueuedAttributes])

	s.password = &password.Case{
		Transfer: s.transfer,
	}

	s.SUT = &password.Consumer{
		Password: s.password,
	}
}

func (s *PasswordTestSuite) TestConsumer() {
	event := messages.RandomWithAttributes(new(events.UserResetQueuedAttributes), true)

	s.transfer.Mock.On("Submit", event.Attributes)

	s.NoError(s.SUT.On(event))

	s.transfer.Mock.AssertExpectations(s.T())
}

func TestUnitPasswordSuite(t *testing.T) {
	suite.Run(t, new(PasswordTestSuite))
}
