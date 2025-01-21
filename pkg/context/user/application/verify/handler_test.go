package verify_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type VerifyTestSuite struct {
	suite.Suite
	SUT        roles.CommandHandler
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

	aggregate := user.RandomPrimitive()

	aggregate.Verified = user.VerifiedWithFalseValue()

	id, err := user.NewID(attributes.ID)

	s.NoError(err)

	aggregate.ID = id

	verify, err := user.NewID(attributes.Verify)

	s.NoError(err)

	aggregate.Verify = verify

	criteria := &user.Criteria{
		ID: id,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	registered := *aggregate

	registered.Verified = user.VerifiedWithTrueValue()

	registered.Verify = nil

	s.repository.Mock.On("Update", &registered)

	command := messages.RandomWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
