package communications

import (
	"fmt"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
)

type QueryBusSuite struct {
	suite.Suite
	SUT     queries.Bus
	Handler *QueryHandlerMock
}

func (s *QueryBusSuite) TestRegister() {
	s.NoError(s.SUT.Register(messages.Random[queries.Query]().Key, s.Handler))
}

func (s *QueryBusSuite) TestRegisterErrDuplicateCommand() {
	key := messages.Random[queries.Query]().Key

	s.NoError(s.SUT.Register(key, s.Handler))

	err := s.SUT.Register(key, s.Handler)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Register",
		What:  fmt.Sprintf("%s already registered", key),
		Why: errors.Meta{
			"Query": key,
		},
	}}

	s.Equal(expected, actual)
}

func (s *QueryBusSuite) TestAsk() {
	query := messages.Random[queries.Query]()

	s.NoError(s.SUT.Register(query.Key, s.Handler))

	response := messages.Random[queries.Response]()

	s.Handler.Mock.On("Handle", query).Return(response)

	actual, err := s.SUT.Ask(query)

	s.NoError(err)

	s.Handler.Mock.AssertExpectations(s.T())

	expected := response

	s.Equal(expected, actual)
}

func (s *QueryBusSuite) TestAskErrMissingHandler() {
	query := messages.Random[queries.Query]()

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
