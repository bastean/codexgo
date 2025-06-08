package login_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type LoginTestSuite struct {
	suite.Frozen
	SUT        roles.QueryHandler
	login      *login.Case
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
		Case: s.login,
	}
}

func (s *LoginTestSuite) TestHandle() {
	aggregate := user.Mother().UserValidFromPrimitive()

	plainPassword := user.Mother().PlainPasswordValid()

	criteria := &user.Criteria{
		Email: aggregate.Email,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	s.hasher.Mock.On("Compare", aggregate.Password.Value(), plainPassword.Value())

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
		Password: plainPassword.Value(),
	}

	query := messages.Mother().MessageValidWithAttributes(attributes, false)

	actual, err := s.SUT.Handle(query)

	s.NoError(err)

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())

	s.Equal(expected, actual)
}

func (s *LoginTestSuite) TestHandleErrMissingRequired() {
	plainPassword := user.Mother().PlainPasswordValid()

	attributes := &login.QueryAttributes{
		Password: plainPassword.Value(),
	}

	query := messages.Mother().MessageValidWithAttributes(attributes, false)

	_, err := s.SUT.Handle(query)

	var actual *errors.Failure

	s.ErrorAs(err, &actual)

	expected := &errors.Failure{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "login/*Case/Run",
		What:  "Email or Username required",
	}}

	s.Equal(expected, actual)
}

func TestUnitLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
