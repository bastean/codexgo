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
	SUT        commands.Handler
	verify     cases.Verify
	repository *persistence.RepositoryMock
}

func (s *VerifyTestSuite) SetupTest() {
	s.repository = new(persistence.RepositoryMock)

	s.verify = &verify.Case{
		Repository: s.repository,
	}

	s.SUT = &verify.Handler{
		Verify: s.verify,
	}
}

func (s *VerifyTestSuite) TestHandle() {
	attributes := verify.CommandRandomAttributes()

	registered := user.RandomPrimitive()

	registered.Verified = user.VerifiedWithFalseValue()

	id, err := user.NewID(attributes.ID)

	s.NoError(err)

	registered.ID = id

	verify, err := user.NewID(attributes.Verify)

	s.NoError(err)

	registered.Verify = verify

	criteria := &repository.Criteria{
		ID: id,
	}

	s.repository.Mock.On("Search", criteria).Return(registered)

	aggregate, err := user.FromPrimitive(registered.ToPrimitive())

	s.NoError(err)

	aggregate.Verified = user.VerifiedWithTrueValue()

	aggregate.Verify = nil

	s.repository.Mock.On("Update", aggregate)

	command := messages.RandomWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
