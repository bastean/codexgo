package verify_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type VerifyTestSuite struct {
	suite.Suite
	sut        commands.Handler
	verify     cases.Verify
	repository *persistence.RepositoryMock
}

func (s *VerifyTestSuite) SetupTest() {
	s.repository = new(persistence.RepositoryMock)

	s.verify = &verify.Case{
		Repository: s.repository,
	}

	s.sut = &verify.Handler{
		Verify: s.verify,
	}
}

func (s *VerifyTestSuite) TestHandle() {
	attributes := verify.CommandRandomAttributes()

	aggregate := user.RandomPrimitive()

	id, err := user.NewID(attributes.ID)

	s.NoError(err)

	aggregate.ID = id

	criteria := &repository.Criteria{
		ID: id,
	}

	s.repository.On("Search", criteria).Return(aggregate)

	s.repository.On("Verify", id)

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	s.NoError(s.sut.Handle(command))

	s.repository.AssertExpectations(s.T())
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
