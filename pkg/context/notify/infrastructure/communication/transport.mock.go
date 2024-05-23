package communication

import (
	"github.com/stretchr/testify/mock"
)

type TransportMock struct {
	mock.Mock
}

func (transport *TransportMock) Submit(data any) error {
	transport.Called(data)
	return nil
}
