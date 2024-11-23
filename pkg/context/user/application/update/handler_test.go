package update_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type UpdateTestSuite struct {
	suite.Suite
	sut        commands.Handler
	update     cases.Update
	hashing    *ciphers.HashingMock
	repository *persistence.UserMock
}

func (s *UpdateTestSuite) SetupTest() {
	s.repository = new(persistence.UserMock)

	s.hashing = new(ciphers.HashingMock)

	s.update = &update.Case{
		Repository: s.repository,
		Hashing:    s.hashing,
	}

	s.sut = &update.Handler{
		Update: s.update,
	}
}

func (s *UpdateTestSuite) TestHandle() {
	attributes := update.CommandRandomAttributes()

	aggregate, err := user.FromRaw(&user.Primitive{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.Password,
	})

	s.NoError(err)

	criteria := &repository.SearchCriteria{
		ID: aggregate.ID,
	}

	hashed := user.CipherPasswordWithValidValue()

	aggregate.CipherPassword = hashed

	s.repository.On("Search", criteria).Return(aggregate)

	s.hashing.On("IsNotEqual", aggregate.CipherPassword.Value, attributes.Password).Return(false)

	s.hashing.On("Hash", attributes.UpdatedPassword).Return(hashed.Value)

	s.repository.On("Update", aggregate)

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	s.NoError(s.sut.Handle(command))

	s.repository.AssertExpectations(s.T())

	s.hashing.AssertExpectations(s.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
