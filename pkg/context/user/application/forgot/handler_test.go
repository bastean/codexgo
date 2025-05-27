package forgot_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/forgot"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type ForgotTestSuite struct {
	suite.Frozen
	SUT        roles.CommandHandler
	forgot     *forgot.Case
	repository *persistence.RepositoryMock
	bus        *communications.EventBusMock
}

func (s *ForgotTestSuite) SetupSuite() {
	s.bus = new(communications.EventBusMock)

	s.repository = new(persistence.RepositoryMock)

	s.forgot = &forgot.Case{
		Repository: s.repository,
		EventBus:   s.bus,
	}

	s.SUT = &forgot.Handler{
		Case: s.forgot,
	}
}

func (s *ForgotTestSuite) TestHandle() {
	attributes := forgot.Mother().CommandAttributesValid()

	aggregate := user.Mother().UserValidFromPrimitive("ResetToken")

	aggregate.Email = values.Mother().EmailNew(attributes.Email)

	criteria := &user.Criteria{
		Email: aggregate.Email,
	}

	s.repository.Mock.On("Search", criteria).
		Run(func(args mock.Arguments) {
			s.SetTimeAfter(time.Hour)
		}).
		Return(aggregate)

	aggregate = user.Mother().UserCopy(aggregate)

	resetToken := values.Mother().TokenNew(attributes.ResetToken)

	aggregate.ResetToken = resetToken

	s.SetTimeAfter(time.Hour)

	s.NoError(aggregate.UpdatedStamp())

	s.UnsetTimeAfter()

	s.repository.Mock.On("Update", aggregate)

	s.bus.Mock.On("Publish", messages.New(
		user.ResetQueuedKey,
		&user.ResetQueuedAttributes{
			ResetToken: resetToken.Value(),
			ID:         aggregate.ID.Value(),
			Email:      aggregate.Email.Value(),
			Username:   aggregate.Username.Value(),
		},
		new(user.ResetQueuedMeta),
	))

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.bus.Mock.AssertExpectations(s.T())
}

func TestUnitForgotSuite(t *testing.T) {
	suite.Run(t, new(ForgotTestSuite))
}
