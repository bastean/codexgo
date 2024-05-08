package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	commandMother "github.com/bastean/codexgo/pkg/context/user/application/delete/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	cryptographicMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic/mock"
	persistenceMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/mock"
	"github.com/stretchr/testify/suite"
)

type UserDeleteTestSuite struct {
	suite.Suite
	sut        *delete.CommandHandler
	delete     *delete.Delete
	hashing    *cryptographicMock.HashingMock
	repository *persistenceMock.RepositoryMock
}

func (suite *UserDeleteTestSuite) SetupTest() {
	suite.repository = new(persistenceMock.RepositoryMock)
	suite.hashing = new(cryptographicMock.HashingMock)
	suite.delete = &delete.Delete{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}
	suite.sut = &delete.CommandHandler{
		UseCase: suite.delete,
	}
}

func (suite *UserDeleteTestSuite) TestDelete() {
	command := commandMother.Random()

	userId, _ := valueObject.NewId(command.Id)

	suite.repository.On("Delete", userId)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitUserDeleteSuite(t *testing.T) {
	suite.Run(t, new(UserDeleteTestSuite))
}
