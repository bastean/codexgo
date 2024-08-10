package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/delete"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type DeleteTestSuite struct {
	suite.Suite
	sut        handlers.Command[*delete.Command]
	delete     usecase.Delete
	hashing    *cryptographic.HashingMock
	repository *persistence.UserMock
}

func (suite *DeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.delete = &delete.Delete{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &delete.Handler{
		Delete: suite.delete,
	}
}

func (suite *DeleteTestSuite) TestDelete() {
	random := user.Random()

	command := &delete.Command{
		Id:       random.Id.Value,
		Password: random.Password.Value,
	}

	criteria := &repository.SearchCriteria{
		Id: random.Id,
	}

	suite.repository.On("Search", criteria).Return(random)

	suite.hashing.On("IsNotEqual", random.Password.Value, random.Password.Value).Return(false)

	suite.repository.On("Delete", random.Id)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitDeleteSuite(t *testing.T) {
	suite.Run(t, new(DeleteTestSuite))
}
