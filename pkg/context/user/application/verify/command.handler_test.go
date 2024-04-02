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
	suite.repository = persistenceMock.NewRepositoryMock()
	suite.verify = verify.NewVerify(suite.repository)
	suite.sut = verify.NewCommandHandler(suite.verify)
}

func (suite *UserVerifyTestSuite) TestVerify() {
	command := commandMother.Random()

	user := aggregateMother.Random()

	idVO := valueObject.NewId(command.Id)

	user.Id = idVO

	user.Password = nil

	filter := model.RepositorySearchFilter{Id: idVO}

	suite.repository.On("Search", filter).Return(user)

	verifiedVO := valueObject.NewVerified(true)

	user.Verified = verifiedVO

	suite.repository.On("Update", user)

	suite.sut.Handle(command)

	suite.repository.AssertExpectations(suite.T())
}

func TestUserVerifySuite(t *testing.T) {
	suite.Run(t, new(UserVerifyTestSuite))
}
