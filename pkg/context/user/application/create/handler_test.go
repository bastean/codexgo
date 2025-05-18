package create_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type CreateTestSuite struct {
	suite.Frozen
	SUT        roles.CommandHandler
	create     *create.Case
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
	bus        *communications.EventBusMock
}

func (s *CreateTestSuite) SetupSuite() {
	s.bus = new(communications.EventBusMock)

	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.create = &create.Case{
		Hasher:     s.hasher,
		Repository: s.repository,
		EventBus:   s.bus,
	}

	s.SUT = &create.Handler{
		Case: s.create,
	}
}

func (s *CreateTestSuite) TestHandle() {
	attributes := create.Mother().CommandAttributesValid()

	hashed := user.Mother().PasswordValid()

	s.hasher.Mock.On("Hash", attributes.Password).Return(hashed.Value())

	aggregate, err := user.New(&user.Required{
		VerifyToken: attributes.VerifyToken,
		ID:          attributes.ID,
		Email:       attributes.Email,
		Username:    attributes.Username,
		Password:    hashed.Value(),
	})

	s.NoError(err)

	s.repository.Mock.On("Create", aggregate)

	for _, event := range aggregate.Events {
		s.bus.Mock.On("Publish", event)
	}

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.hasher.Mock.AssertExpectations(s.T())

	s.repository.Mock.AssertExpectations(s.T())

	s.bus.Mock.AssertExpectations(s.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
