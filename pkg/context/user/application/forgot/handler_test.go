package forgot_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/events"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/forgot"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type ForgotTestSuite struct {
	suite.Suite
	SUT        roles.CommandHandler
	forgot     cases.Forgot
	repository *persistence.RepositoryMock
	bus        *communications.EventBusMock
}

func (s *ForgotTestSuite) SetupSuite() {
	s.bus = new(communications.EventBusMock)

	s.repository = new(persistence.RepositoryMock)

	s.forgot = &forgot.Case{
		Repository: s.repository,
	}

	s.SUT = &forgot.Handler{
		Forgot:   s.forgot,
		EventBus: s.bus,
	}
}

func (s *ForgotTestSuite) SetupTest() {
	s.NoError(os.Setenv("GOTEST_FROZEN", "1"))
}

func (s *ForgotTestSuite) TestHandle() {
	attributes := forgot.Mother.CommandValidAttributes()

	email, err := values.New[*user.Email](attributes.Email)

	s.NoError(err)

	aggregate := user.Mother.UserValid()

	aggregate.Email = email

	criteria := &user.Criteria{
		Email: aggregate.Email,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	reset, err := values.New[*user.ID](attributes.Reset)

	s.NoError(err)

	registered := *aggregate

	registered.Reset = reset

	s.repository.Mock.On("Update", &registered)

	registered.Record(messages.New(
		events.UserResetQueuedKey,
		&events.UserResetQueuedAttributes{
			Reset:    registered.Reset.Value(),
			ID:       registered.ID.Value(),
			Email:    registered.Email.Value(),
			Username: registered.Username.Value(),
		},
		new(events.UserResetQueuedMeta),
	))

	for _, event := range registered.Events {
		s.bus.Mock.On("Publish", event)
	}

	command := messages.Mother.MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.bus.Mock.AssertExpectations(s.T())
}

func (s *ForgotTestSuite) TearDownTest() {
	s.NoError(os.Unsetenv("GOTEST_FROZEN"))
}

func TestUnitForgotSuite(t *testing.T) {
	suite.Run(t, new(ForgotTestSuite))
}
