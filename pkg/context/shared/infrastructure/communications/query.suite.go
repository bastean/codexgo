package communications

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/errors"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/messages"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/roles"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/suite"
)

type QueryBusSuite struct {
	suite.Default
	SUT     roles.QueryBus
	Handler *QueryHandlerMock
	Query   *messages.Message
}

func (s *QueryBusSuite) SetupTest() {
	s.Query = messages.Mother().MessageValid()
}

func (s *QueryBusSuite) TestRegister() {
	s.NoError(s.SUT.Register(s.Query.Key, s.Handler))
}

func (s *QueryBusSuite) TestRegisterErrDuplicateQuery() {
	s.NoError(s.SUT.Register(s.Query.Key, s.Handler))

	err := s.SUT.Register(s.Query.Key, s.Handler)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Register",
		What:  "Already registered",
		Why: errors.Meta{
			"Key": s.Query.Key.Value(),
		},
	}}

	s.Equal(expected, actual)
}

func (s *QueryBusSuite) TestAsk() {
	s.NoError(s.SUT.Register(s.Query.Key, s.Handler))

	response := messages.Mother().MessageValid()

	s.Handler.Mock.On("Handle", s.Query).Return(response)

	actual, err := s.SUT.Ask(s.Query)

	s.NoError(err)

	s.Handler.Mock.AssertExpectations(s.T())

	expected := response

	s.Equal(expected, actual)
}

func (s *QueryBusSuite) TestAskErrMissingHandler() {
	s.Query = messages.Mother().MessageValid()

	_, err := s.SUT.Ask(s.Query)

	var actual *errors.Internal

	s.ErrorAs(err, &actual)

	expected := &errors.Internal{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "Ask",
		What:  "Failure to execute a Query without a Handler",
		Why: errors.Meta{
			"ID":  s.Query.ID.Value(),
			"Key": s.Query.Key.Value(),
		},
	}}

	s.Equal(expected, actual)
}
