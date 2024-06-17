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

type UpdateHandlerTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*update.Command]
	usecase    models.UseCase[*update.Input, types.Empty]
	hashing    *cryptographic.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *UpdateHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.usecase = &update.Update{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &update.Handler{
		UseCase: suite.usecase,
	}
}

func (suite *UpdateHandlerTestSuite) TestUpdate() {
	command := update.RandomCommand()

	user, _ := aggregate.NewUser(&aggregate.UserPrimitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.UpdatedPassword,
	})

	idVO, _ := valueobj.NewId(command.Id)

	criteria := &model.RepositorySearchCriteria{
		Id: idVO,
	}

	suite.repository.On("Search", criteria).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value(), command.Password).Return(false)

	suite.repository.On("Update", user)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())
}

func TestUnitUpdateHandlerSuite(t *testing.T) {
	suite.Run(t, new(UpdateHandlerTestSuite))
}
