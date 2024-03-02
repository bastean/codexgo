package cryptographic

import "github.com/stretchr/testify/mock"

type HashingMock struct {
	mock.Mock
}

func (m *HashingMock) Hash(plain string) string {
	return plain
}

func (m *HashingMock) IsNotEqual(hashed, plain string) bool {
	return false
}
