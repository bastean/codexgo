package memory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type QueryBusTestSuite struct {
	suite.Suite
	sut     queries.Bus
	handler *communications.QueryHandlerMock
}

func (s *QueryBusTestSuite) SetupTest() {
	s.handler = new(communications.QueryHandlerMock)

	s.sut = &memory.QueryBus{
		Handlers: make(map[queries.Key]queries.Handler),
	}
}

func (s *QueryBusTestSuite) TestRegister() {
	s.NoError(s.sut.Register(messages.Random[queries.Query]().Key, s.handler))
}

func (s *QueryBusTestSuite) TestRegisterErrDuplicateCommand() {
	key := messages.Random[queries.Query]().Key

	s.NoError(s.sut.Register(key, s.handler))

	err := s.sut.Register(key, s.handler)

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

func (s *QueryBusTestSuite) TestAsk() {
	query := messages.Random[queries.Query]()

	s.NoError(s.sut.Register(query.Key, s.handler))

	response := messages.Random[queries.Response]()

	s.handler.On("Handle", query).Return(response)

	actual, err := s.sut.Ask(query)

	s.NoError(err)

	s.handler.AssertExpectations(s.T())

	expected := response

	s.Equal(expected, actual)
}

func (s *QueryBusTestSuite) TestAskErrMissingHandler() {
	query := messages.Random[queries.Query]()

	_, err := s.sut.Ask(query)

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

func TestIntegrationQueryBusSuite(t *testing.T) {
	suite.Run(t, new(QueryBusTestSuite))
}
