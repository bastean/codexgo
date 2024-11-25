package delete_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type DeleteTestSuite struct {
	suite.Suite
	sut        commands.Handler
	delete     cases.Delete
	hashing    *ciphers.HashingMock
	repository *persistence.RepositoryMock
}

func (s *DeleteTestSuite) SetupTest() {
	s.repository = new(persistence.RepositoryMock)

	s.hashing = new(ciphers.HashingMock)

	s.delete = &delete.Case{
		Repository: s.repository,
		Hashing:    s.hashing,
	}

	s.sut = &delete.Handler{
		Delete: s.delete,
	}
}

func (s *DeleteTestSuite) TestHandle() {
	aggregate := user.RandomPrimitive()

	plain := user.PlainPasswordWithValidValue()

	criteria := &repository.SearchCriteria{
		ID: aggregate.ID,
	}

	s.repository.On("Search", criteria).Return(aggregate)

	s.hashing.On("IsNotEqual", aggregate.CipherPassword.Value, plain.Value).Return(false)

	s.repository.On("Delete", aggregate.ID)

	attributes := &delete.CommandAttributes{
		ID:       aggregate.ID.Value,
		Password: plain.Value,
	}

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	s.NoError(s.sut.Handle(command))

	s.repository.AssertExpectations(s.T())

	s.hashing.AssertExpectations(s.T())
}

func TestUnitDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
