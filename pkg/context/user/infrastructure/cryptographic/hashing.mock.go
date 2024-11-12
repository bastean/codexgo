package cryptographic

import (
	"github.com/stretchr/testify/mock"
)

type HashingMock struct {
	mock.Mock
}

func (m *HashingMock) Hash(plain string) (string, error) {
	args := m.Called(plain)
	return args.Get(0).(string), nil
}

func (m *HashingMock) IsNotEqual(hashed, plain string) bool {
	args := m.Called(hashed, plain)
	return args.Get(0).(bool)
}
