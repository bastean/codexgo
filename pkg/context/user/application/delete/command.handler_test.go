package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/smodel"
	"github.com/bastean/codexgo/pkg/context/shared/domain/stype"
	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UserDeleteTestSuite struct {
	suite.Suite
	sut        smodel.CommandHandler[*delete.Command]
	useCase    smodel.UseCase[smodel.ValueObject[string], *stype.Empty]
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
	command := delete.RandomCommand()

	userId, _ := valueobj.NewId(command.Id)

	suite.repository.On("Delete", userId)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitUserDeleteSuite(t *testing.T) {
	suite.Run(t, new(UserDeleteTestSuite))
}
