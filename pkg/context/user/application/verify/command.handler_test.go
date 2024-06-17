package verify_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/models"
	"github.com/bastean/codexgo/pkg/context/shared/domain/types"
	"github.com/bastean/codexgo/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/valueobj"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type VerifyHandlerTestSuite struct {
	suite.Suite
	sut        models.CommandHandler[*verify.Command]
	usecase    models.UseCase[models.ValueObject[string], types.Empty]
	repository *persistence.RepositoryMock
}

func (suite *VerifyHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.usecase = &verify.Verify{
		Repository: suite.repository,
	}

	suite.sut = &verify.Handler{
		UseCase: suite.usecase,
	}
}

func (suite *VerifyHandlerTestSuite) TestVerify() {
	command := verify.RandomCommand()

	user := aggregate.RandomUser()

	idVO, _ := valueobj.NewId(command.Id)

	user.Id = idVO

	criteria := &model.RepositorySearchCriteria{
		Id: idVO,
	}

	suite.repository.On("Search", criteria).Return(user)

	suite.repository.On("Verify", idVO)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitVerifyHandlerSuite(t *testing.T) {
	suite.Run(t, new(VerifyHandlerTestSuite))
}
