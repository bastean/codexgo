package verify_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
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

	aggregate.ID = values.Mother().IDNew(attributes.ID)

	aggregate.VerifyToken = values.Mother().TokenNew(attributes.VerifyToken)

	aggregate.Verified = user.Mother().VerifiedFalse()

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).
		Run(func(args mock.Arguments) {
			s.SetTimeAfter(time.Hour)
		}).
		Return(aggregate)

	aggregate = user.Mother().UserCopy(aggregate)

	verified := user.Mother().VerifiedTrue()

	verified.SetUpdated(time.Now().Add(time.Hour))

	aggregate.Verified = verified

	aggregate.VerifyToken = nil

	s.SetTimeAfter(time.Hour)

	s.NoError(aggregate.UpdatedStamp())

	s.UnsetTimeAfter()

	s.repository.Mock.On("Update", aggregate)

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
