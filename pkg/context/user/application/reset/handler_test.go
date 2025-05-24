package reset_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/reset"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type ResetTestSuite struct {
	suite.Frozen
	SUT        roles.CommandHandler
	reset      *reset.Case
	repository *persistence.RepositoryMock
	hasher     *ciphers.HasherMock
}

func (s *ResetTestSuite) SetupSuite() {
	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.reset = &reset.Case{
		Repository: s.repository,
		Hasher:     s.hasher,
	}

	s.SUT = &reset.Handler{
		Case: s.reset,
	}
}

func (s *ResetTestSuite) TestHandle() {
	attributes := reset.Mother().CommandAttributesValid()

	aggregate := user.Mother().UserValidFromPrimitive()

	aggregate.ResetToken = user.Mother().IDNew(attributes.ResetToken)

	aggregate.ID = user.Mother().IDNew(attributes.ID)

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	hashed := user.Mother().PasswordValid()

	s.hasher.Mock.On("Hash", attributes.Password).
		Run(func(args mock.Arguments) {
			s.SetTimeAfter(12)
		}).
		Return(hashed.Value())

	aggregate = user.Mother().UserCopy(aggregate)

	hashed.SetUpdated(time.Now().Add(12))

	aggregate.Password = hashed

	aggregate.ResetToken = nil

	s.SetTimeAfter(12)

	s.NoError(aggregate.UpdatedStamp())

	s.UnsetTimeAfter()

	s.repository.Mock.On("Update", aggregate)

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())
}

func TestUnitResetSuite(t *testing.T) {
	suite.Run(t, new(ResetTestSuite))
}
