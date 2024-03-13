package update_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	commandMother "github.com/bastean/codexgo/test/pkg/context/user/application/update"
	valueObjectMother "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
	cryptographicMock "github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/cryptographic"
	persistenceMock "github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/persistence"
	"github.com/stretchr/testify/suite"
)

type UserUpdateTestSuite struct {
	suite.Suite
	sut        *update.CommandHandler
	update     *update.Update
	hashing    *cryptographicMock.HashingMock
	repository *persistenceMock.RepositoryMock
}

func (suite *UserUpdateTestSuite) SetupTest() {
	suite.repository = persistenceMock.NewRepositoryMock()
	suite.hashing = cryptographicMock.NewHashingMock()
	suite.update = update.NewUpdate(suite.repository, suite.hashing)
	suite.sut = update.NewCommandHandler(suite.update)
}

func (suite *UserUpdateTestSuite) TestUpdate() {
	command := commandMother.Random()

	user := aggregate.Create(command.Id, command.Email, command.Username, command.Password)

	idVO := valueObjectMother.NewId(command.Id)

	filter := model.RepositorySearchFilter{Id: idVO}

	suite.repository.On("Search", filter).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value, command.Password).Return(false)

	suite.repository.On("Update", user)

	suite.sut.Handle(command)

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())
}

func TestUserUpdateSuite(t *testing.T) {
	suite.Run(t, new(UserUpdateTestSuite))
}
