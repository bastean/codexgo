package ciphers

import (
	"github.com/bastean/codexgo/v4/pkg/context/shared/domain/services/mock"
)

type HasherMock struct {
	mock.Default
}

func (m *HasherMock) Hash(plain string) (string, error) {
	args := m.Called(plain)
	return args.Get(0).(string), nil
}

func (m *HasherMock) Compare(hashed, plain string) error {
	m.Called(hashed, plain)
	return nil
}
