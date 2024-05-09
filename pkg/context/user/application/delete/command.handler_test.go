package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	commandMother "github.com/bastean/codexgo/pkg/context/user/application/delete/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	cryptographicMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic/mock"
	persistenceMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/mock"
	"github.com/stretchr/testify/suite"
)

type UserDeleteTestSuite struct {
	suite.Suite
	sut        model.CommandHandler[*delete.Command]
	useCase    model.UseCase[model.ValueObject[string], *types.Empty]
	hashing    *cryptographicMock.HashingMock
	repository *persistenceMock.RepositoryMock
}

func (suite *UserDeleteTestSuite) SetupTest() {
	suite.repository = new(persistenceMock.RepositoryMock)
	suite.hashing = new(cryptographicMock.HashingMock)
	suite.useCase = &delete.Delete{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}
	suite.sut = &delete.CommandHandler{
		UseCase: suite.useCase,
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
