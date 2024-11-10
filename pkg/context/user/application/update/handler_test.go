package update_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/update"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type UpdateTestSuite struct {
	suite.Suite
	sut        commands.Handler
	update     cases.Update
	hashing    *cryptographic.HashingMock
	repository *persistence.UserMock
}

func (suite *UpdateTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.update = &update.Update{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &update.Handler{
		Update: suite.update,
	}
}

func (suite *UpdateTestSuite) TestHandle() {
	attributes := update.CommandRandomAttributes()

	account, err := user.FromPrimitive(&user.Primitive{
		Id:       attributes.Id,
		Email:    attributes.Email,
		Username: attributes.Username,
		Password: attributes.UpdatedPassword,
	})

	suite.NoError(err)

	id, err := user.NewId(attributes.Id)

	suite.NoError(err)

	criteria := &repository.SearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(account)

	suite.hashing.On("IsNotEqual", account.Password.Value, attributes.Password).Return(false)

	suite.repository.On("Update", account)

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())
}

func TestUnitUpdateSuite(t *testing.T) {
	suite.Run(t, new(UpdateTestSuite))
}
