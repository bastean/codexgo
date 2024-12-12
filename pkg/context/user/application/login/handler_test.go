package login_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/ciphers"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
)

type LoginTestSuite struct {
	suite.Suite
	SUT        queries.Handler
	login      cases.Login
	hasher     *ciphers.HasherMock
	repository *persistence.RepositoryMock
}

func (s *LoginTestSuite) SetupTest() {
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

func (s *LoginTestSuite) TestHandle() {
	aggregate := user.RandomPrimitive()

	plain := user.PlainPasswordWithValidValue()

	criteria := &repository.Criteria{
		Email: aggregate.Email,
	}

	s.repository.Mock.On("Search", criteria).Return(aggregate)

	s.hasher.Mock.On("IsNotEqual", aggregate.CipherPassword.Value, plain.Value).Return(false)

	expected := messages.New[queries.Response](
		login.ResponseKey,
		(*login.ResponseAttributes)(aggregate.ToPrimitive()),
		new(login.ResponseMeta),
	)

	attributes := &login.QueryAttributes{
		Email:    aggregate.Email.Value,
		Password: plain.Value,
	}

	query := messages.RandomWithAttributes[queries.Query](attributes, false)

	actual, err := s.SUT.Handle(query)

	s.NoError(err)

	s.repository.Mock.AssertExpectations(s.T())

	s.hasher.Mock.AssertExpectations(s.T())

	s.EqualValues(expected, actual)
}

func (s *LoginTestSuite) TestHandleErrMissingRequired() {
	plain := user.PlainPasswordWithValidValue()

	attributes := &login.QueryAttributes{
		Password: plain.Value,
	}

	query := messages.RandomWithAttributes[queries.Query](attributes, false)

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

func TestUnitLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
