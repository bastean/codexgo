package memory_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/query"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications"
	"github.com/bastean/codexgo/v4/pkg/context/shared/infrastructure/communications/memory"
)

type QueryBusTestSuite struct {
	suite.Suite
	sut      query.Bus
	query    *communications.QueryMock
	response *communications.ResponseMock
	handler  *communications.QueryHandlerMock
}

func (suite *QueryBusTestSuite) SetupTest() {
	suite.query = new(communications.QueryMock)

	suite.response = new(communications.ResponseMock)

	suite.handler = new(communications.QueryHandlerMock)

	suite.sut = &memory.QueryBus{
		Handlers: make(map[query.Type]query.Handler),
	}
}

func (suite *QueryBusTestSuite) TestRegister() {
	const ask query.Type = "query.testing.register"
	suite.NoError(suite.sut.Register(ask, suite.handler))
}

func (suite *QueryBusTestSuite) TestRegisterErrDuplicateCommand() {
	const ask query.Type = "query.testing.register_duplicate"

	suite.NoError(suite.sut.Register(ask, suite.handler))

	err := suite.sut.Register(ask, suite.handler)

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Register",
		What:  fmt.Sprintf("%s already registered", ask),
		Why: errors.Meta{
			"Query": ask,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *QueryBusTestSuite) TestAsk() {
	const ask query.Type = "query.testing.ask"

	suite.NoError(suite.sut.Register(ask, suite.handler))

	suite.query.On("Type").Return(ask)

	suite.handler.On("Handle", suite.query).Return(suite.response)

	actual, err := suite.sut.Ask(suite.query)

	suite.NoError(err)

	suite.query.AssertExpectations(suite.T())

	suite.handler.AssertExpectations(suite.T())

	expected := suite.response

	suite.Equal(expected, actual)
}

func (suite *QueryBusTestSuite) TestAskErrMissingHandler() {
	const ask query.Type = "query.testing.ask_missing"

	suite.query.On("Type").Return(ask)

	_, err := suite.sut.Ask(suite.query)

	suite.query.AssertExpectations(suite.T())

	var actual *errors.Internal

	suite.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Ask",
		What:  "Failure to execute a Query without a Handler",
		Why: errors.Meta{
			"Query": ask,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestIntegrationQueryBusSuite(t *testing.T) {
	suite.Run(t, new(QueryBusTestSuite))
}
