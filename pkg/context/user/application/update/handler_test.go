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
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
}

func (s *UpdateTestSuite) SetupTest() {
	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.update = &update.Case{
		Repository: s.repository,
		Hasher:     s.hasher,
	}

	s.sut = &update.Handler{
		Update: s.update,
	}
}

func (s *UpdateTestSuite) TestHandle() {
	registered := user.RandomPrimitive()

	attributes := update.CommandRandomAttributes()

	attributes.ID = registered.ID.Value

	aggregate, err := user.FromRaw(&user.Primitive{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.Password,
	})

	s.NoError(err)

	criteria := &repository.Criteria{
		ID: aggregate.ID,
	}

	s.repository.On("Search", criteria).Return(registered)

	s.hasher.On("IsNotEqual", registered.CipherPassword.Value, aggregate.PlainPassword.Value).Return(false)

	hashed := user.CipherPasswordWithValidValue()

	s.hasher.On("Hash", attributes.UpdatedPassword).Return(hashed.Value)

	aggregate.CipherPassword = hashed

	s.NoError(err)

	aggregate.Verified = registered.Verified

	s.repository.On("Update", aggregate)

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	s.NoError(s.sut.Handle(command))

	s.repository.AssertExpectations(s.T())

	s.hasher.AssertExpectations(s.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
