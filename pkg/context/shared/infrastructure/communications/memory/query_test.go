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

func (suite *QueryBusTestSuite) SetupTest() {
	suite.handler = new(communications.QueryHandlerMock)

	suite.sut = &memory.QueryBus{
		Handlers: make(map[queries.Key]queries.Handler),
	}
}

func (suite *QueryBusTestSuite) TestRegister() {
	suite.NoError(suite.sut.Register(messages.Random[queries.Query]().Key, suite.handler))
}

func (suite *QueryBusTestSuite) TestRegisterErrDuplicateCommand() {
	key := messages.Random[queries.Query]().Key

	suite.NoError(suite.sut.Register(key, suite.handler))

	err := suite.sut.Register(key, suite.handler)

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Register",
		What:  fmt.Sprintf("%s already registered", key),
		Why: errors.Meta{
			"Query": key,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *QueryBusTestSuite) TestAsk() {
	query := messages.Random[queries.Query]()

	suite.NoError(suite.sut.Register(query.Key, suite.handler))

	response := messages.Random[queries.Response]()

	suite.handler.On("Handle", query).Return(response)

	actual, err := suite.sut.Ask(query)

	suite.NoError(err)

	suite.handler.AssertExpectations(suite.T())

	expected := response

	suite.Equal(expected, actual)
}

func (suite *QueryBusTestSuite) TestAskErrMissingHandler() {
	query := messages.Random[queries.Query]()

	_, err := suite.sut.Ask(query)

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Ask",
		What:  "Failure to execute a Query without a Handler",
		Why: errors.Meta{
			"Query": query.Key,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationQueryBusSuite(t *testing.T) {
	suite.Run(t, new(QueryBusTestSuite))
}
