package verify_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
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

func (s *VerifyTestSuite) SetupSuite() {
	s.repository = new(persistence.RepositoryMock)

	s.verify = &verify.Case{
		Repository: s.repository,
	}

	s.SUT = &verify.Handler{
		Verify: s.verify,
	}
}

func (s *VerifyTestSuite) SetupTest() {
	s.NoError(os.Setenv("GOTEST_FROZEN", "1"))
}

func (s *VerifyTestSuite) TestHandle() {
	attributes := verify.Mother.CommandValidAttributes()

	aggregate := user.Mother.UserValidPrimitive()

	aggregate.Verified = user.Mother.VerifiedFalse()

	id, err := values.New[*user.ID](attributes.ID)

	s.NoError(err)

	aggregate.ID = id

	verify, err := values.New[*user.ID](attributes.Verify)

	s.NoError(err)

	aggregate.Verify = verify

	criteria := &user.Criteria{
		ID: id,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	registered := *aggregate

	registered.Verified = user.Mother.VerifiedTrue()

	registered.Verify = nil

	s.repository.Mock.On("Update", &registered)

	command := messages.Mother.MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())
}

func (s *VerifyTestSuite) TearDownTest() {
	s.NoError(os.Unsetenv("GOTEST_FROZEN"))
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
