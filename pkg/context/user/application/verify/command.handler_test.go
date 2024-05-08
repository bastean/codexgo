package verify_test

import (
	"testing"

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
	sut        *verify.CommandHandler
	verify     *verify.Verify
	repository *persistenceMock.RepositoryMock
}

func (suite *UserVerifyTestSuite) SetupTest() {
	suite.repository = new(persistenceMock.RepositoryMock)
	suite.verify = &verify.Verify{
		Repository: suite.repository,
	}
	suite.sut = &verify.CommandHandler{
		UseCase: suite.verify,
	}
}

func (suite *UserVerifyTestSuite) TestVerify() {
	command := commandMother.Random()

	user := aggregateMother.Random()

	idVO, _ := valueObject.NewId(command.Id)

	user.Id = idVO

	user.Password = nil

	filter := model.RepositorySearchCriteria{Id: idVO}

	suite.repository.On("Search", filter).Return(user)

	verifiedVO, _ := valueObject.NewVerified(true)

	user.Verified = verifiedVO

	suite.repository.On("Update", user)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitUserVerifySuite(t *testing.T) {
	suite.Run(t, new(UserVerifyTestSuite))
}
