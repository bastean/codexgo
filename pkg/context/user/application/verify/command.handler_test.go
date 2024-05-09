package verify_test

import (
	"testing"

	sharedModel "github.com/bastean/codexgo/pkg/context/shared/domain/model"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	commandMother "github.com/bastean/codexgo/pkg/context/user/application/verify/mother"
	aggregateMother "github.com/bastean/codexgo/pkg/context/user/domain/aggregate/mother"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueObject"
	persistenceMock "github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence/mock"
	"github.com/stretchr/testify/suite"
)

type UserVerifyTestSuite struct {
	suite.Suite
	sut        sharedModel.CommandHandler[*verify.Command]
	useCase    sharedModel.UseCase[sharedModel.ValueObject[string], *types.Empty]
	repository *persistenceMock.RepositoryMock
}

func (suite *UserVerifyTestSuite) SetupTest() {
	suite.repository = new(persistenceMock.RepositoryMock)
	suite.useCase = &verify.Verify{
		Repository: suite.repository,
	}
	suite.sut = &verify.CommandHandler{
		UseCase: suite.useCase,
	}
}

func (suite *UserVerifyTestSuite) TestVerify() {
	command := commandMother.Random()

	user := aggregateMother.Random()

	idVO, _ := valueObject.NewId(command.Id)

	user.Id = idVO

	user.Password = nil

	filter := model.RepositorySearchCriteria{
		Id: idVO,
	}

	suite.repository.On("Search", filter).Return(user)

	suite.repository.On("Update", user)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitUserVerifySuite(t *testing.T) {
	suite.Run(t, new(UserVerifyTestSuite))
}
