package communications

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/queries"
)

type QueryHandlerMock struct {
	mock.Mock
}

func (m *QueryHandlerMock) Handle(query *queries.Query) (*queries.Response, error) {
	args := m.Called(query)
	return args.Get(0).(*queries.Response), nil
}

type QueryBusMock struct {
	mock.Mock
}

func (m *QueryBusMock) Register(key queries.Key, handler queries.Handler) error {
	m.Called(key, handler)
	return nil
}

func (m *QueryBusMock) Ask(query *queries.Query) (*queries.Response, error) {
	args := m.Called(query)
	return args.Get(0).(*queries.Response), nil
}
