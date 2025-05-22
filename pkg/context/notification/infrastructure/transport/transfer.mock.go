package transport

import (
	"github.com/bastean/codexgo/v4/pkg/context/notification/domain/aggregate/recipient"
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
)

type TransferMock struct {
	mock.Default
}

func (m *TransferMock) Submit(recipient *recipient.Recipient) error {
	m.Called(recipient)
	return nil
}
