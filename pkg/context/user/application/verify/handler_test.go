package verify_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/commands"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/verify"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type VerifyTestSuite struct {
	suite.Suite
	sut        commands.Handler
	verify     cases.Verify
	repository *persistence.UserMock
}

func (suite *VerifyTestSuite) SetupTest() {
	suite.repository = new(persistence.UserMock)

	suite.verify = &verify.Verify{
		Repository: suite.repository,
	}

	suite.sut = &verify.Handler{
		Verify: suite.verify,
	}
}

func (suite *VerifyTestSuite) TestHandle() {
	attributes := verify.CommandRandomAttributes()

	random := user.Random()

	id, err := user.NewId(attributes.Id)

	suite.NoError(err)

	random.Id = id

	criteria := &repository.SearchCriteria{
		Id: id,
	}

	suite.repository.On("Search", criteria).Return(random)

	suite.repository.On("Verify", id)

	command := messages.RandomWithAttributes[commands.Command](attributes, false)

	suite.NoError(suite.sut.Handle(command))

	suite.repository.AssertExpectations(suite.T())
}

func TestUnitVerifySuite(t *testing.T) {
	suite.Run(t, new(VerifyTestSuite))
}