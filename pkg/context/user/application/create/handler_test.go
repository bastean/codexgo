package create_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type CreateTestSuite struct {
	suite.Suite
	SUT        commands.Handler
	create     cases.Create
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
	bus        *communications.EventBusMock
}

func (s *CreateTestSuite) SetupTest() {
	s.bus = new(communications.EventBusMock)

	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.create = &create.Case{
		Hasher:     s.hasher,
		Repository: s.repository,
	}

	s.SUT = &create.Handler{
		Create: s.create,
		Bus:    s.bus,
	}
}

func (s *CreateTestSuite) TestHandle() {
	attributes := create.CommandRandomAttributes()

	aggregate, err := user.FromRaw(&user.Primitive{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.Password,
	})

	s.NoError(err)

	hashed := user.CipherPasswordWithValidValue()

	s.hasher.Mock.On("Hash", aggregate.PlainPassword.Value).Return(hashed.Value)

	aggregate.CipherPassword = hashed

	s.repository.Mock.On("Create", aggregate)

	for _, event := range aggregate.Events {
		s.bus.Mock.On("Publish", event)
	}

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.bus.Mock.AssertExpectations(s.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
