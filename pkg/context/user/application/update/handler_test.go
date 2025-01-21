package update_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type UpdateTestSuite struct {
	suite.Suite
	SUT        roles.CommandHandler
	update     cases.Update
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
}

func (s *UpdateTestSuite) SetupTest() {
	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.update = &update.Case{
		Repository: s.repository,
		Hasher:     s.hasher,
	}

	s.SUT = &update.Handler{
		Update: s.update,
	}
}

func (s *UpdateTestSuite) TestHandle() {
	registered := user.RandomPrimitive()

	attributes := update.CommandRandomAttributes()

	attributes.ID = registered.ID.Value

	aggregate, err := user.FromRaw(&user.Primitive{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.Password,
	})

	s.NoError(err)

	criteria := &user.Criteria{
		ID: aggregate.ID,
	}

	s.repository.Mock.On("Search", criteria).Return(registered)

	s.hasher.Mock.On("Compare", registered.CipherPassword.Value, aggregate.PlainPassword.Value)

	hashed := user.CipherPasswordWithValidValue()

	s.hasher.Mock.On("Hash", attributes.UpdatedPassword).Return(hashed.Value)

	aggregate.CipherPassword = hashed

	aggregate.Created = registered.Created
	aggregate.Updated = registered.Updated
	aggregate.Verified = registered.Verified

	s.repository.Mock.On("Update", aggregate)

	command := messages.RandomWithAttributes(attributes, false)

	s.NoError(s.SUT.Handle(command))

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
