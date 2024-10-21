package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/query"
)

type QueryMock struct {
	mock.Mock
}

func (m *QueryMock) Type() query.Type {
	args := m.Called()
	return args.Get(0).(query.Type)
}

type ResponseMock struct {
	mock.Mock
}

func (m *ResponseMock) Type() query.Type {
	args := m.Called()
	return args.Get(0).(query.Type)
}

type QueryHandlerMock struct {
	mock.Mock
}

func (m *QueryHandlerMock) SubscribedTo() query.Type {
	args := m.Called()
	return args.Get(0).(query.Type)
}

func (m *QueryHandlerMock) ReplyTo() query.Type {
	args := m.Called()
	return args.Get(0).(query.Type)
}

func (m *QueryHandlerMock) Handle(ask query.Query) (query.Response, error) {
	args := m.Called(ask)
	return args.Get(0).(query.Response), nil
}

type QueryBusMock struct {
	mock.Mock
}

func (m *QueryBusMock) Register(ask query.Type, handler query.Handler) error {
	m.Called(ask, handler)
	return nil
}

func (m *QueryBusMock) Ask(ask query.Query) (query.Response, error) {
	args := m.Called(ask)
	return args.Get(0).(query.Response), nil
}
