package delete_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type DeleteTestSuite struct {
	suite.Suite
	sut        commands.Handler
	delete     cases.Delete
	hashing    *cryptographic.HashingMock
	repository *persistence.UserMock
}

func (suite *DeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.delete = &delete.Delete{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &delete.Handler{
		Delete: suite.delete,
	}
}

func (suite *DeleteTestSuite) TestHandle() {
	account := user.Random()

	criteria := &repository.SearchCriteria{
		Id: account.Id,
	}

	suite.repository.On("Search", criteria).Return(account)

	suite.hashing.On("IsNotEqual", account.Password.Value, account.Password.Value).Return(false)

	suite.repository.On("Delete", account.Id)

	attributes := &delete.CommandAttributes{
		Id:       account.Id.Value,
		Password: account.Password.Value,
	}

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())
}

func TestUnitDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
