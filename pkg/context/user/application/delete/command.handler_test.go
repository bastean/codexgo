package delete_test

import (
	"testing"

	sharedValueObject "github.com/bastean/codexgo/pkg/context/shared/domain/valueObject"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	commandMother "github.com/bastean/codexgo/pkg/context/user/application/delete/mother"
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
	suite.repository = persistenceMock.NewRepositoryMock()
	suite.hashing = cryptographicMock.NewHashingMock()
	suite.delete = delete.NewDelete(suite.repository, suite.hashing)
	suite.sut = delete.NewCommandHandler(suite.delete)
}

func (suite *UserDeleteTestSuite) TestDelete() {
	command := commandMother.Random()

	userId := sharedValueObject.NewId(command.Id)

	suite.repository.On("Delete", userId)

	suite.sut.Handle(command)

	suite.repository.AssertExpectations(suite.T())
}

func TestUserDeleteSuite(t *testing.T) {
	suite.Run(t, new(UserDeleteTestSuite))
}
