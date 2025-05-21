package transport

import (
	"github.com/stretchr/testify/mock"

	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
)

type TransferMock struct {
	mock.Mock
}

func (m *TransferMock) Submit(recipient *recipient.Recipient) error {
	m.Called(recipient)
	return nil
}
