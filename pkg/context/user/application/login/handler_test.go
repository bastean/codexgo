package login_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type LoginTestSuite struct {
	suite.Suite
	SUT        roles.QueryHandler
	login      cases.Login
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
}

func (s *LoginTestSuite) SetupSuite() {
	s.repository = new(persistence.RepositoryMock)

	s.hasher = new(ciphers.HasherMock)

	s.login = &login.Case{
		Repository: s.repository,
		Hasher:     s.hasher,
	}

	s.SUT = &login.Handler{
		Login: s.login,
	}
}

func (s *LoginTestSuite) SetupTest() {
	s.NoError(os.Setenv("GOTEST_FROZEN", "1"))
}

func (s *LoginTestSuite) TestHandle() {
	aggregate := user.Mother.UserValidPrimitive()

	plain := user.Mother.PlainPasswordValid()

	criteria := &user.Criteria{
		Email: aggregate.Email,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	s.hasher.Mock.On("Compare", aggregate.CipherPassword.Value(), plain.Value())

	response := &login.ResponseAttributes{
		ID:       aggregate.ID.Value(),
		Email:    aggregate.Email.Value(),
		Username: aggregate.Username.Value(),
		Verified: aggregate.Verified.Value(),
	}

	expected := messages.New(
		login.ResponseKey,
		response,
		new(login.ResponseMeta),
	)

	attributes := &login.QueryAttributes{
		Email:    aggregate.Email.Value(),
		Password: plain.Value(),
	}

	query := messages.Mother.MessageValidWithAttributes(attributes, false)

	actual, err := s.SUT.Handle(query)

	s.NoError(err)

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())

	s.EqualValues(expected, actual)
}

func (s *LoginTestSuite) TestHandleErrMissingRequired() {
	plain := user.Mother.PlainPasswordValid()

	attributes := &login.QueryAttributes{
		Password: plain.Value(),
	}

	query := messages.Mother.MessageValidWithAttributes(attributes, false)

	_, err := s.SUT.Handle(query)

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Handle",
		What:  "Email or Username required",
	}}

	s.Equal(expected, actual)
}

func (s *LoginTestSuite) TearDownTest() {
	s.NoError(os.Unsetenv("GOTEST_FROZEN"))
}

func TestUnitLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
