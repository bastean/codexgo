package update_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserUpdateTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*update.Command]
	useCase    models.UseCase[*update.Command, *types.Empty]
	hashing    *cryptographic.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *UserUpdateTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.useCase = &update.Update{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &update.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *UserUpdateTestSuite) TestUpdate() {
	command := update.RandomCommand()

	user, _ := aggregate.NewUser(command.Id, command.Email, command.Username, command.Password)

	idVO, _ := valueobj.NewId(command.Id)

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
