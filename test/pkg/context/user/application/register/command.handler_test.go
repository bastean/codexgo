package register_test

import (
	"testing"

	"github.com/bastean/codexgo/pkg/context/user/application/register"
	command "github.com/bastean/codexgo/test/pkg/context/user/application/register"
	"github.com/bastean/codexgo/test/pkg/context/user/domain/aggregate"
	"github.com/bastean/codexgo/test/pkg/context/user/infrastructure/mock/persistence"
	"github.com/stretchr/testify/suite"
)

type UserRegisterTestSuite struct {
	suite.Suite
	repository *persistence.RepositoryMock
	register   *register.Register
	handler    *register.CommandHandler
}

func (suite *UserRegisterTestSuite) SetupTest() {
	suite.repository = new(persistence.RepositoryMock)
	suite.register = &register.Register{Repository: suite.repository}
	suite.handler = &register.CommandHandler{Register: *suite.register}
}

func (suite *UserRegisterTestSuite) TestRegister() {
	command := command.Random()

	user := aggregate.FromCommand(*command)

	suite.repository.On("Save", user)

	suite.handler.Handle(*command)

	suite.repository.AssertCalled(suite.T(), "Save", user)
}

func TestUserRegisterSuite(t *testing.T) {
	suite.Run(t, new(UserRegisterTestSuite))
}
