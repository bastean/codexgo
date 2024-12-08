package transport

import (
	"github.com/stretchr/testify/mock"
)

type TransferMock[T any] struct {
	mock.Mock
}

func (m *TransferMock[T]) Submit(data T) error {
	m.Called(data)
	return nil
}
