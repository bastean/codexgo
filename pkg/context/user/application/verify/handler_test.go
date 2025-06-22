package verify_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/aggregates/token"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type VerifyTestSuite struct {
	suite.Frozen
	SUT        roles.CommandHandler
	verify     *verify.Case
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
}

func (s *VerifyTestSuite) SetupSuite() {
	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.verify = &verify.Case{
		Repository: s.repository,
		Hasher:     s.hasher,
	}

	s.SUT = &verify.Handler{
		Case: s.verify,
	}
}

func (s *VerifyTestSuite) TestHandle() {
	attributes := verify.Mother().CommandAttributesValid()

	aggregate := user.Mother().UserValidFromPrimitive()

	aggregate.ID = values.Mother().IDNew(attributes.ID)

	aggregate.VerifyToken = token.Mother().TokenNew(attributes.VerifyToken)

	aggregate.Verified = user.Mother().VerifiedFalse()

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	s.hasher.Mock.On("Compare", aggregate.Password.Value(), attributes.Password)

	aggregate = user.Mother().UserCopy(aggregate)

	aggregate.Verified = user.Mother().VerifiedReplace(aggregate.Verified, true)

	aggregate.VerifyToken = nil

	s.NoError(aggregate.UpdatedStamp())

	s.repository.Mock.On("Update", aggregate)

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
