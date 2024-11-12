package create_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/create"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type CreateTestSuite struct {
	suite.Suite
	sut        commands.Handler
	create     cases.Create
	repository *persistence.UserMock
	bus        *communications.EventBusMock
}

func (s *CreateTestSuite) SetupTest() {
	s.bus = new(communications.EventBusMock)

	s.repository = new(persistence.UserMock)

	s.create = &create.Case{
		Repository: s.repository,
	}

	s.sut = &create.Handler{
		Create: s.create,
		Bus:    s.bus,
	}
}

func (s *CreateTestSuite) TestHandle() {
	attributes := create.CommandRandomAttributes()

	account, err := user.New(&user.Primitive{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.Password,
	})

	s.NoError(err)

	s.repository.On("Create", account)

	for _, event := range account.Events {
		s.bus.On("Publish", event)
	}

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	s.NoError(s.sut.Handle(command))

	s.repository.AssertExpectations(s.T())

	s.bus.AssertExpectations(s.T())
}

func TestUnitCreateSuite(t *testing.T) {
	suite.Run(t, new(CreateTestSuite))
}
