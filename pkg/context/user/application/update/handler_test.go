package update_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/values"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type UpdateTestSuite struct {
	suite.Frozen
	SUT        roles.CommandHandler
	update     *update.Case
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
}

func (s *UpdateTestSuite) SetupSuite() {
	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.update = &update.Case{
		Repository: s.repository,
		Hasher:     s.hasher,
	}

	s.SUT = &update.Handler{
		Case: s.update,
	}
}

func (s *UpdateTestSuite) TestHandle() {
	aggregate := user.Mother().UserValidFromPrimitive()

	attributes := update.Mother().CommandAttributesValid()

	attributes.ID = aggregate.ID.Value()

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	s.hasher.Mock.On("Compare", aggregate.Password.Value(), attributes.Password)

	hashed := user.Mother().PasswordValid().Value()

	s.hasher.Mock.On("Hash", attributes.UpdatedPassword).Return(hashed)

	aggregate = user.Mother().UserCopy(aggregate)

	aggregate.Password = user.Mother().PasswordReplace(aggregate.Password, hashed)

	aggregate.Email = values.Mother().EmailReplace(aggregate.Email, attributes.Email)

	aggregate.Username = values.Mother().UsernameReplace(aggregate.Username, attributes.Username)

	s.NoError(aggregate.UpdatedStamp())

	s.repository.Mock.On("Update", aggregate)

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
