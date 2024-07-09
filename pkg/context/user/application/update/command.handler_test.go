package update_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/shared/domain/handlers"
	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/pkg/context/user/domain/model"
	"github.com/bastean/codexgo/pkg/context/user/domain/usecase"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type UpdateHandlerTestSuite struct {
	suite.Suite
	sut        handlers.Command[*update.Command]
	update     usecase.Update
	hashing    *cryptographic.HashingMock
	repository *persistence.RepositoryMock
}

func (suite *UpdateHandlerTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)

	suite.hashing = new(cryptographic.HashingMock)

	suite.update = &update.Update{
		Repository: suite.repository,
		Hashing:    suite.hashing,
	}

	suite.sut = &update.Handler{
		Update: suite.update,
	}
}

func (suite *UpdateHandlerTestSuite) TestUpdate() {
	command := update.RandomCommand()

	new, _ := user.New(&user.Primitive{
		Id:       command.Id,
		Email:    command.Email,
		Username: command.Username,
		Password: command.UpdatedPassword,
	})

	id, _ := user.NewId(command.Id)

	criteria := &model.RepositorySearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(new)

	suite.hashing.On("IsNotEqual", new.Password.Value, command.Password).Return(false)

	suite.repository.On("Update", new)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())

	suite.hashing.AssertExpectations(suite.T())
}

func TestUnitUpdateHandlerSuite(t *testing.T) {
	suite.Run(t, new(UpdateHandlerTestSuite))
}
