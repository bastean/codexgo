package ciphers

import (
	"github.com/stretchr/testify/mock"
)

type HasherMock struct {
	mock.Mock
}

func (m *HasherMock) Hash(plain string) (string, error) {
	args := m.Called(plain)
	return args.Get(0).(string), nil
}

func (m *HasherMock) IsNotEqual(hashed, plain string) bool {
	args := m.Called(hashed, plain)
	return args.Get(0).(bool)
}
