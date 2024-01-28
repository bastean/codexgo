package update_test

import (
	"testing"

	"github.com/bastean/codexgo/context/pkg/user/application/update"
	"github.com/bastean/codexgo/context/pkg/user/domain/aggregate"
	"github.com/bastean/codexgo/context/pkg/user/domain/repository"
	"github.com/bastean/codexgo/test/contexts/crud/user/__mocks__/infrastructure/cryptographic"
	"github.com/bastean/codexgo/test/contexts/crud/user/__mocks__/infrastructure/persistence"
	command "github.com/bastean/codexgo/test/contexts/crud/user/application/update"
	create "github.com/bastean/codexgo/test/contexts/crud/user/domain/valueObject"
	"github.com/stretchr/testify/suite"
)

type UserUpdateTestSuite struct {
	suite.Suite
	repository *persistence.UserRepositoryMock
	hashing    *cryptographic.UserHashingMock
	update     *update.Update
	handler    *update.CommandHandler
}

func (suite *UserUpdateTestSuite) SetupTest() {
	suite.repository = new(persistence.UserRepositoryMock)
	suite.hashing = new(cryptographic.UserHashingMock)
	suite.update = &update.Update{Repository: suite.repository, Hashing: suite.hashing}
	suite.handler = &update.CommandHandler{Update: *suite.update}
}

func (suite *UserUpdateTestSuite) TestUpdate() {
	command := command.Random()

	user := aggregate.Create(command.Id, command.Email, command.Username, command.UpdatedPassword)

	idVO := create.NewId(command.Id)

	filter := repository.Filter{Id: idVO}

	suite.repository.On("Search", filter).Return(user)

	suite.repository.On("Update", user)

	suite.handler.Handle(*command)

	suite.repository.AssertCalled(suite.T(), "Search", filter)

	suite.repository.AssertCalled(suite.T(), "Update", user)
}

func TestUserUpdateSuite(t *testing.T) {
	suite.Run(t, new(UserUpdateTestSuite))
}
