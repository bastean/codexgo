package login_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

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
	sut        queries.Handler
	login      cases.Login
	hashing    *ciphers.HashingMock
	repository *persistence.UserMock
}

func (s *LoginTestSuite) SetupTest() {
	s.repository = new(persistence.UserMock)

	s.hashing = new(ciphers.HashingMock)

	s.login = &login.Case{
		Repository: s.repository,
		Hashing:    s.hashing,
	}

	s.sut = &login.Handler{
		Login: s.login,
	}
}

func (s *LoginTestSuite) TestHandle() {
	aggregate := user.RandomPrimitive()

	plain := user.PlainPasswordWithValidValue()

	criteria := &repository.SearchCriteria{
		Email: aggregate.Email,
	}

	s.repository.On("Search", criteria).Return(aggregate)

	s.hashing.On("IsNotEqual", aggregate.CipherPassword.Value, plain.Value).Return(false)

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

	actual, err := s.sut.Handle(query)

	s.NoError(err)

	s.repository.AssertExpectations(s.T())

	s.hashing.AssertExpectations(s.T())

	s.EqualValues(expected, actual)
}

func TestUnitLoginSuite(t *testing.T) {
	suite.Run(t, new(LoginTestSuite))
}
