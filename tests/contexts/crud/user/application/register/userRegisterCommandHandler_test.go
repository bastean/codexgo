package register_test

import (
	"testing"

	"github.com/bastean/codexgo/context/pkg/user/application/register"
	"github.com/bastean/codexgo/test/contexts/crud/user/__mocks__/infrastructure/persistence"
	command "github.com/bastean/codexgo/test/contexts/crud/user/application/register"
	"github.com/bastean/codexgo/test/contexts/crud/user/domain/aggregate"
	"github.com/stretchr/testify/suite"
)

type UserRegisterTestSuite struct {
	suite.Suite
	repository *persistence.UserRepositoryMock
	register   *register.Register
	handler    *register.CommandHandler
}

func (suite *UserRegisterTestSuite) SetupTest() {
	suite.repository = new(persistence.UserRepositoryMock)
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
