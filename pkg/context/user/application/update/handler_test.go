package update_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UpdateTestSuite struct {
	suite.Suite
	sut        commands.Handler
	update     cases.Update
	hashing    *cryptographic.HashingMock
	repository *persistence.UserMock
}

func (s *UpdateTestSuite) SetupTest() {
	s.repository = new(persistence.UserMock)

	s.hashing = new(cryptographic.HashingMock)

	s.update = &update.Case{
		Repository: s.repository,
		Hashing:    s.hashing,
	}

	s.sut = &update.Handler{
		Update: s.update,
	}
}

func (s *UpdateTestSuite) TestHandle() {
	attributes := update.CommandRandomAttributes()

	account, err := user.FromPrimitive(&user.Primitive{
		ID:       attributes.ID,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.UpdatedPassword,
	})

	s.NoError(err)

	id, err := user.NewID(attributes.ID)

	s.NoError(err)

	criteria := &repository.SearchCriteria{
		ID: id,
	}

	s.repository.On("Search", criteria).Return(account)

	s.hashing.On("IsNotEqual", account.Password.Value, attributes.Password).Return(false)

	s.repository.On("Update", account)

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	s.NoError(s.sut.Handle(command))

	s.repository.AssertExpectations(s.T())

	s.hashing.AssertExpectations(s.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
