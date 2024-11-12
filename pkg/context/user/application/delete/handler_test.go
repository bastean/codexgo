package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type DeleteTestSuite struct {
	suite.Suite
	sut        commands.Handler
	delete     cases.Delete
	hashing    *cryptographic.HashingMock
	repository *persistence.UserMock
}

func (s *DeleteTestSuite) SetupTest() {
	s.repository = new(persistence.UserMock)

	s.hashing = new(cryptographic.HashingMock)

	s.delete = &delete.Case{
		Repository: s.repository,
		Hashing:    s.hashing,
	}

	s.sut = &delete.Handler{
		Delete: s.delete,
	}
}

func (s *DeleteTestSuite) TestHandle() {
	account := user.Random()

	criteria := &repository.SearchCriteria{
		ID: account.ID,
	}

	s.repository.On("Search", criteria).Return(account)

	s.hashing.On("IsNotEqual", account.Password.Value, account.Password.Value).Return(false)

	s.repository.On("Delete", account.ID)

	attributes := &delete.CommandAttributes{
		ID:       account.ID.Value,
		Password: account.Password.Value,
	}

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	s.NoError(s.sut.Handle(command))

	s.repository.AssertExpectations(s.T())

	s.hashing.AssertExpectations(s.T())
}

func TestUnitDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
