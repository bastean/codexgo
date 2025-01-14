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

func (m *HasherMock) Compare(hashed, plain string) error {
	m.Called(hashed, plain)
	return nil
}
