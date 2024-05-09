package update_test

import (
	"testing"

	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	commandMother "github.com/bastean/codexgo/pkg/context/user/application/update/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	cryptographicMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic/mock"
	persistenceMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/mock"
	"github.com/stretchr/testify/suite"
)

type UserUpdateTestSuite struct {
	suite.Suite
	sut        sharedModel.CommandHandler[*update.Command]
	useCase    sharedModel.UseCase[*update.Command, *types.Empty]
	hashing    *cryptographicMock.HashingMock
	repository *persistenceMock.RepositoryMock
}

func (suite *UserUpdateTestSuite) SetupTest() {
	suite.repository = new(persistenceMock.RepositoryMock)
	suite.hashing = new(cryptographicMock.HashingMock)
	suite.useCase = &update.Update{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}
	suite.sut = &update.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *UserUpdateTestSuite) TestUpdate() {
	command := commandMother.Random()

	user, _ := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	idVO, _ := valueObject.NewId(command.Id)

	filter := model.RepositorySearchCriteria{
		Id: idVO,
	}

	suite.repository.On("Search", filter).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value(), command.Password).Return(false)

	suite.repository.On("Update", user)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())
}

func TestUnitUserUpdateSuite(t *testing.T) {
	suite.Run(t, new(UserUpdateTestSuite))
}
