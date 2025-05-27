package update_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/time"
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

	s.hasher.Mock.On("Compare", aggregate.Password.Value(), attributes.Password).
		Run(func(args mock.Arguments) {
			s.SetTimeAfter(time.Hour)
		})

	hashed := user.Mother().PasswordValid()

	s.hasher.Mock.On("Hash", attributes.UpdatedPassword).Return(hashed.Value())

	aggregate = user.Mother().UserCopy(aggregate)

	email := values.Mother().EmailNew(attributes.Email)

	email.SetUpdated(time.Now().Add(time.Hour))

	aggregate.Email = email

	username := values.Mother().UsernameNew(attributes.Username)

	username.SetUpdated(time.Now().Add(time.Hour))

	aggregate.Username = username

	hashed.SetUpdated(time.Now().Add(time.Hour))

	aggregate.Password = hashed

	s.SetTimeAfter(time.Hour)

	s.NoError(aggregate.UpdatedStamp())

	s.UnsetTimeAfter()

	s.repository.Mock.On("Update", aggregate)

	command := messages.Mother().MessageValidWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
