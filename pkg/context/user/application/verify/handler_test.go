package verify_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type VerifyTestSuite struct {
	suite.Frozen
	SUT        roles.CommandHandler
	verify     *verify.Case
	repository *persistence.RepositoryMock
}

func (s *VerifyTestSuite) SetupSuite() {
	s.repository = new(persistence.RepositoryMock)

	s.verify = &verify.Case{
		Repository: s.repository,
	}

	s.SUT = &verify.Handler{
		Case: s.verify,
	}
}

func (s *VerifyTestSuite) TestHandle() {
	attributes := verify.Mother().CommandAttributesValid()

	aggregate := user.Mother().UserValidFromPrimitive()

	aggregate.ID = user.Mother().IDNew(attributes.ID)

	aggregate.VerifyToken = user.Mother().IDNew(attributes.VerifyToken)

	aggregate.Verified = user.Mother().VerifiedFalse()

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).
		Run(func(args suite.Arguments) {
			s.SetTimeAfter(12)
		}).
		Return(aggregate)

	aggregateWithVerify := *aggregate

	verified := user.Mother().VerifiedTrue()

	verified.SetUpdated(time.Now().Add(12))

	aggregateWithVerify.Verified = verified

	aggregateWithVerify.VerifyToken = nil

	s.repository.Mock.On("Update", &aggregateWithVerify)

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
