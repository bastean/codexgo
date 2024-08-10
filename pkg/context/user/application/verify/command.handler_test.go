package verify_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type VerifyTestSuite struct {
	suite.Suite
	sut        handlers.Command[*verify.Command]
	verify     usecase.Verify
	repository *persistence.UserMock
}

func (suite *VerifyTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.verify = &verify.Verify{
		Repository: suite.repository,
	}

	suite.sut = &verify.Handler{
		Verify: suite.verify,
	}
}

func (suite *VerifyTestSuite) TestVerify() {
	command := verify.RandomCommand()

	random := user.Random()

	id, err := user.NewId(command.Id)

	suite.NoError(err)

	random.Id = id

	criteria := &repository.SearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(random)

	suite.repository.On("Verify", id)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}
