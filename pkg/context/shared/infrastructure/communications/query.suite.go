package communications

import (
	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
)

type QueryBusSuite struct {
	suite.Suite
	SUT     roles.QueryBus
	Handler *QueryHandlerMock
}

func (s *QueryBusSuite) TestRegister() {
	s.NoError(s.SUT.Register(messages.Mother.MessageValid().Key, s.Handler))
}

func (s *QueryBusSuite) TestRegisterErrDuplicateQuery() {
	key := messages.Mother.MessageValid().Key

	s.NoError(s.SUT.Register(key, s.Handler))

	err := s.SUT.Register(key, s.Handler)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Register",
		What:  key.Value() + " already registered",
		Why: errors.Meta{
			"Query": key,
		},
	}}

	s.Equal(expected, actual)
}

func (s *QueryBusSuite) TestAsk() {
	query := messages.Mother.MessageValid()

	s.NoError(s.SUT.Register(query.Key, s.Handler))

	response := messages.Mother.MessageValid()

	s.Handler.Mock.On("Handle", query).Return(response)

	actual, err := s.SUT.Ask(query)

	s.NoError(err)

	s.Handler.Mock.AssertExpectations(s.T())

	expected := response

	s.Equal(expected, actual)
}

func (s *QueryBusSuite) TestAskErrMissingHandler() {
	query := messages.Mother.MessageValid()

	_, err := s.SUT.Ask(query)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Ask",
		What:  "Failure to execute a Query without a Handler",
		Why: errors.Meta{
			"Query": query.Key,
		},
	}}

	s.Equal(expected, actual)
}
