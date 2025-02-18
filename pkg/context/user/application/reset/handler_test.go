package reset_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/reset"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type ResetTestSuite struct {
	suite.Suite
	SUT        roles.CommandHandler
	reset      cases.Reset
	repository *persistence.RepositoryMock
	hasher     *ciphers.HasherMock
}

func (s *ResetTestSuite) SetupTest() {
	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.reset = &reset.Case{
		Repository: s.repository,
		Hasher:     s.hasher,
	}

	s.SUT = &reset.Handler{
		Reset: s.reset,
	}
}

func (s *ResetTestSuite) TestHandle() {
	attributes := reset.CommandRandomAttributes()

	reset, err := user.NewID(attributes.Reset)

	s.NoError(err)

	id, err := user.NewID(attributes.ID)

	s.NoError(err)

	aggregate := user.Random()

	aggregate.Reset = reset

	aggregate.ID = id

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	hashed := user.CipherPasswordWithValidValue()

	s.hasher.Mock.On("Hash", attributes.Password).Return(hashed.Value)

	registered := *aggregate

	registered.CipherPassword = hashed

	registered.Reset = nil

	s.repository.Mock.On("Update", &registered)

	command := messages.RandomWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())
}

func TestUnitResetSuite(t *testing.T) {
	suite.Run(t, new(ResetTestSuite))
}
