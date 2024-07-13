package transports

import (
	"github.com/stretchr/testify/mock"
)

type TransferMock struct {
	mock.Mock
}

func (transfer *TransferMock) Submit(data any) error {
	transfer.Called(data)
	return nil
}
