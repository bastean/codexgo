package login_test

import (
	"testing"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/user/application/login"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/aggregate/user"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/cases"
	"github.com/bastean/codexgo/v4/pkg/context/user/domain/repository"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/cryptographic"
	"github.com/bastean/codexgo/v4/pkg/context/user/infrastructure/persistence"
	"github.com/stretchr/testify/suite"
)

type LoginTestSuite struct {
	suite.Suite
	sut        queries.Handler
	login      cases.Login
	hashing    *cryptographic.HashingMock
	repository *persistence.UserMock
}

func (s *LoginTestSuite) SetupTest() {
	s.repository = new(persistence.UserMock)

	s.hashing = new(cryptographic.HashingMock)

	s.login = &login.Case{
		Repository: s.repository,
		Hashing:    s.hashing,
	}

	s.sut = &login.Handler{
		Login: s.login,
	}
}

func (s *LoginTestSuite) TestHandle() {
	account := user.Random()

	criteria := &repository.SearchCriteria{
		Email: account.Email,
	}

	s.repository.On("Search", criteria).Return(account)

	s.hashing.On("IsNotEqual", account.Password.Value, account.Password.Value).Return(false)

	expected := messages.New[queries.Response](
		login.ResponseKey,
		(*login.ResponseAttributes)(account.ToPrimitive()),
		new(login.ResponseMeta),
	)

	attributes := &login.QueryAttributes{
		Email:    account.Email.Value,
		Password: account.Password.Value,
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
