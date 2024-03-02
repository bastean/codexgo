package delete_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/application/delete"
	command "github.com/bastean/codexgo/test/pkg/context/user/application/delete"
	create "github.com/bastean/codexgo/test/pkg/context/user/domain/valueObject"
	"github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/cryptographic"
	"github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/persistence"
	"github.com/stretchr/testify/suite"
)

type UserDeleteTestSuite struct {
	suite.Suite
	repository *persistence.RepositoryMock
	hashing    *cryptographic.HashingMock
	delete     *delete.Delete
	handler    *delete.CommandHandler
}

func (suite *UserDeleteTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
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
