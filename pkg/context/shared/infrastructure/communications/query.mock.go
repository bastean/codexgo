package communications

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
	"github.com/stretchr/testify/mock"
)

type QueryMock struct {
	mock.Mock
}

func (m *QueryMock) Type() queries.Type {
	args := m.Called()
	return args.Get(0).(queries.Type)
}

type ResponseMock struct {
	mock.Mock
}

func (m *ResponseMock) Type() queries.Type {
	args := m.Called()
	return args.Get(0).(queries.Type)
}

type QueryHandlerMock struct {
	mock.Mock
}

func (m *QueryHandlerMock) SubscribedTo() queries.Type {
	args := m.Called()
	return args.Get(0).(queries.Type)
}

func (m *QueryHandlerMock) ReplyTo() queries.Type {
	args := m.Called()
	return args.Get(0).(queries.Type)
}

func (m *QueryHandlerMock) Handle(ask queries.Query) (queries.Response, error) {
	args := m.Called(ask)
	return args.Get(0).(queries.Response), nil
}

type QueryBusMock struct {
	mock.Mock
}

func (m *QueryBusMock) Register(ask queries.Type, handler queries.Handler) error {
	m.Called(ask, handler)
	return nil
}

func (m *QueryBusMock) Ask(ask queries.Query) (queries.Response, error) {
	args := m.Called(ask)
	return args.Get(0).(queries.Response), nil
}
