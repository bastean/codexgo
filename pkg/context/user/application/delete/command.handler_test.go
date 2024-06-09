package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserDeleteTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*delete.Command]
	useCase    models.UseCase[*delete.Input, types.Empty]
	hashing    *cryptographic.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *UserDeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.useCase = &delete.Delete{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &delete.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *UserDeleteTestSuite) TestDelete() {
	user := aggregate.RandomUser()

	command := &delete.Command{
		Id:       user.Id.Value(),
		Password: user.Password.Value(),
	}

	filter := model.RepositorySearchCriteria{
		Id: user.Id,
	}

	suite.repository.On("Search", filter).Return(user)

	suite.hashing.On("IsNotEqual", user.Password.Value(), user.Password.Value()).Return(false)

	suite.repository.On("Delete", user.Id)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitUserDeleteSuite(t *testing.T) {
	suite.Run(t, new(UserDeleteTestSuite))
}
