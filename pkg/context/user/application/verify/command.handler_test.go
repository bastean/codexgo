package verify_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type VerifyHandlerTestSuite struct {
	suite.Suite
	sut        handlers.Command[*verify.Command]
	verify     usecase.Verify
	repository *persistence.RepositoryMock
}

func (suite *VerifyHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.verify = &verify.Verify{
		Repository: suite.repository,
	}

	suite.sut = &verify.Handler{
		Verify: suite.verify,
	}
}

func (suite *VerifyHandlerTestSuite) TestVerify() {
	command := verify.RandomCommand()

	random := user.Random()

	id, _ := user.NewId(command.Id)

	random.Id = id

	criteria := &model.RepositorySearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(random)

	suite.repository.On("Verify", id)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitVerifyHandlerSuite(t *testing.T) {
	suite.Run(t, new(VerifyHandlerTestSuite))
}
