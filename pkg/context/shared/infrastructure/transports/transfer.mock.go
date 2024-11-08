package transports

import (
	"github.com/stretchr/testify/mock"
)

type TransferMock[T any] struct {
	mock.Mock
}

func (transfer *TransferMock[T]) Submit(data T) error {
	transfer.Called(data)
	return nil
}
