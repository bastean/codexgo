package update_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/application/update"
	"github.com/bastean/codexgo/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/pkg/context/user/domain/repository"
	command "github.com/bastean/codexgo/test/pkg/context/user/application/update"
	create "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
	"github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/cryptographic"
	"github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/persistence"
	"github.com/stretchr/testify/suite"
)

type UserUpdateTestSuite struct {
	suite.Suite
	repository *persistence.RepositoryMock
	hashing    *cryptographic.HashingMock
	update     *update.Update
	handler    *update.CommandHandler
}

func (suite *UserUpdateTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.hashing = new(cryptographic.HashingMock)
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
