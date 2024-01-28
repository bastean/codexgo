package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/context/pkg/user/application/delete"
	"github.com/bastean/codexgo/test/contexts/crud/user/__mocks__/infrastructure/cryptographic"
	"github.com/bastean/codexgo/test/contexts/crud/user/__mocks__/infrastructure/persistence"
	command "github.com/bastean/codexgo/test/contexts/crud/user/application/delete"
	create "github.com/bastean/codexgo/test/contexts/crud/user/domain/valueObject"
	"github.com/stretchr/testify/suite"
)

type UserDeleteTestSuite struct {
	suite.Suite
	repository *persistence.UserRepositoryMock
	hashing    *cryptographic.UserHashingMock
	delete     *delete.Delete
	handler    *delete.CommandHandler
}

func (suite *UserDeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.UserRepositoryMock)
	suite.delete = &delete.Delete{Repository: suite.repository, Hashing: suite.hashing}
	suite.handler = &delete.CommandHandler{Delete: *suite.delete}
}

func (suite *UserDeleteTestSuite) TestDelete() {
	command := command.Random()

	userId := create.NewId(command.Id)

	suite.repository.On("Delete", userId)

	suite.handler.Handle(*command)

	suite.repository.AssertCalled(suite.T(), "Delete", userId)
}

func TestUserDeleteSuite(t *testing.T) {
	suite.Run(t, new(UserDeleteTestSuite))
}
