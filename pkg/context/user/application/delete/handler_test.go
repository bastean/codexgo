package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type DeleteTestSuite struct {
	suite.Frozen
	SUT        roles.CommandHandler
	delete     *delete.Case
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
}

func (s *DeleteTestSuite) SetupSuite() {
	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.delete = &delete.Case{
		Repository: s.repository,
		Hasher:     s.hasher,
	}

	s.SUT = &delete.Handler{
		Case: s.delete,
	}
}

func (s *DeleteTestSuite) TestHandle() {
	aggregate := user.Mother().UserValidFromPrimitive()

	plainPassword := user.Mother().PlainPasswordValid()

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	s.hasher.Mock.On("Compare", aggregate.Password.Value(), plainPassword.Value())

	s.repository.Mock.On("Delete", aggregate.ID)

	attributes := &delete.CommandAttributes{
		ID:       aggregate.ID.Value(),
		Password: plainPassword.Value(),
	}

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())
}

func TestUnitDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
